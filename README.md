# Introduction

This project tries to create a solution for high load demand surges in a taxi as a service provider platform like uber.  
This is a proof of concept and is not intended to be used in production.

# Software architecture

The mind map and the concepts I have put together to create the finanl output.    

We start from the simplest possible thing, `someone needs a cab` and `someone is a cab`. There are many many people demanding a cab and many cabs offering their services. We want a very very very minimal and simple solution, ignoring the distances, destination, user experience and other things.  

Here I have demonstrated a normal not-realtime HTTP based process for requesting a cab.

![simple request response](./docs/images/simple-cycle.png)

Probably, when the cab reaches the pessanger, the system will set the demand status to `in-journey` and when they reach the destination, the system will set the demand status to `completed`, we do not need to achieve these at this time.   
Infact we are going to investigate just step 1, and step 2. The problem is that sometimes demand is so much higher than the number of cabs available. for example we can only service around 100 people in a hour, but demand is beyond that number, also we can conclude that if we are that much popular and there are so much demands, we can somehow make more money.  
The proposed solution to this problem is getting a higher fee for the exact demand, with the exact joureny, so this will lead to less people using our service in short terms, hence flattening the surge in demands in this short time peroid, also it will cause more caps register into our product because of high wage in long terms.   

How to implement it?  
We just need to check how many demands are created in last one hour, and if the number is higher than the threshold, we will multiply the journey cost by the rate specified.  
It is clear that executing a query for each demand to calculate the surge rate, is not practical, so I thought of two approaches: 
1. Cache the number of demands with a short time life span, for example we query the database and cache into redis for one minute, and surge service calculates the rating from this value, after one minute it will be refreshed hence old demands will be removed and new ones will be added, one problems is that this is not so much realtime, if we have a surge of demands in under one minute, we will not be able to detect it.  
2. More realtime aproach, we can cach the number of demands in redis on service startup, and each new demand will increase the cach value by +1, and after one hour, we will decrease the cach value by -1, hence the value is realtime and surge rating is calculated very precisely.   

Aproach one doesn't requrie much resources, but aproach two uses a time queue to decrease the cached values, for example, if we have one milion demands in a time span of one hour(worst case scenario, not going to happen very soon), there will be almost one milion `decrease by one after one hour` messages in the queue, the worst possible thing can happen is that you will need about `one or two` extra GIGs of RAM.  
A small note about the aproach two is that, because there maybe old messages on the queue, the worker that consumes the decrease messages must does the message only if it is created after the timestamp of the cached value, otherwise it will be ignored, to assure the inconsistency.  
Also we must check if one unique user is trying to create many demands in a short time, if so, we will not change the cached demand count.  

![surge service basic](./docs/images/surge-basic.png)

And we need a service to resolve the geographical location of the client into the real life district name or district id, so we are going to open street map, overpass api to obtain these information.  
Also we need a user management service that will handle authentication of our clients.  

![overal design](./docs/images/design.png)

## Resources

Here are some usefull links:  
- [Understand OSM data models.](https://wiki.openstreetmap.org/wiki/Elements)  
- [Tehran OSM relation.](https://www.openstreetmap.org/relation/6663864#map=12/35.7398/51.4933)  
- [Go bee framework.](https://beego.vip/docs/intro/)  
- [Go gin framework.](https://github.com/gin-gonic/gin)  
- [Gorm ORM.](https://github.com/go-gorm/gorm)  
- [Writing a simple API with gin and gorm.](https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/)  
- [Adding swagger support to gin.](https://github.com/swaggo/swag)  
- [Using viper for environment files.](https://github.com/spf13/viper)  
- [Hash user passwords in database.](https://pkg.go.dev/golang.org/x/crypto/bcrypt)  
- [JWT tokens in golang.](https://github.com/golang-jwt/jwt)  
- [Wrting a middleware for gin.](https://sosedoff.com/2014/12/21/gin-middleware.html)  
- [Wrting tests in golang.](https://www.digitalocean.com/community/tutorials/how-to-write-unit-tests-in-go-using-go-test-and-the-testing-package)  

