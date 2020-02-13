# Vacation Finder
Don't know where to go on vacations ? Let the computer decide for you !
## Synopsis
```./main [OPTION]...```
## Description
Sort a random latitude and longitude in function of a given start point and a maximal distance.

```-place PLACE``` define your starting point on a known¹ place (see below). Only accept _strings_.

```-dist DISTANCE``` define the maximal vertical and horizontal distance you want. This information create a square around your starting point in which a random pair of coordinates will be chosen. Only accept _integers_.

```-n NUMBER``` set the number of random points you want to have. Only accept _integers_.

¹ Known places are written in locations.yaml.
## Existing places
- Lyon, FR (45.7597, 4.8342)
- Paris, FR (48.8617, 2.3429)

## Author
Written by ezekiel.

## Copyright
 Copyright © 2020 Free Software Foundation, Inc.  License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>. This is free software: you are free to change and redistribute it. There  is  NO WARRANTY, to the extent permitted by law.