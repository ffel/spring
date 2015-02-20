Spring Calculations [![Build Status](https://travis-ci.org/ffel/spring.svg?branch=master)](https://travis-ci.org/ffel/spring)
===================

See [my physics lab](http://www.myphysicslab.com/spring1.html)

![](spring.png)

    $ go run spring.go > data
    $ gnuplot -p -e "set term svg; set output 'plot.svg'; plot 'data' using 1:2 with linespoints"
