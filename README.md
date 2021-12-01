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
The proposed solution to this problem is getting a higher fee for the exact demand, with the exact joureny, so this will lead to less people using our service, also it will cause more caps register into our product because of high wage.  

## Resources

Here are some usefull links:  
[Understand OSM data models.](https://wiki.openstreetmap.org/wiki/Elements)  
[Tehran OSM relation.](https://www.openstreetmap.org/relation/6663864#map=12/35.7398/51.4933)