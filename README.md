# Vacation Finder
Don't know where to go on vacations ? Let the computer decide for you !
## Synopsis
```./main [OPTION]...```
## Description
Sort a random latitude and longitude in function of a given start point and a maximal distance.

```--place PLACE``` define your starting point on a known place (see below). Only accept _strings_.

```--dist DISTANCE``` define the maximal vertical and horizontal distance you want. This information create a square around your starting point in which a random pair of coordinates will be chosen. Only accept _integers_.

## Existing places
- Lyon, FR (45.7597, 4.8342)