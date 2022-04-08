# go-pi
Calculate PI using Monte [Carlo method](https://en.wikipedia.org/wiki/Monte_Carlo_method).

### Usage
Just run application with number of probes provided in the command line. 

Example: 
```
./go-pi 1000 
```

Application will randomize X coordinates and check if point is inside cirlce. Scope of random numbers is inside of square, that's why we don't have to verify it. 