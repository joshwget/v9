# v9-go

Reimplementation of [v9](https://github.com/joshwget/v9), the successor to Google's v8, written in Go.

Created during [Hack the Planet](http://hacktheplanet.mlh.io/).

## Installation

1. Install Go and [nex](https://crypto.stanford.edu/~blynn/nex/)
2. "make"

## Example

This example will run as expected in v9, and roughly shows what was finished during the weekend.

    var i = 3;
    var j = 3;
    while (i) {
      j = 3;
      while (j) {
        log(j);
        j = j - 1;
      }
      i = i - 1;
    }

    if (true) {
      if ("this string is true") {
        if (5.5) {
          log("this should print");
        }
      }
    }

    if ("") {
      log("should not print");
    }

    var ex = function() {
      log("in a function");
    };

    i = 3;
    while (i) {
      ex();
      i = i - 1;
    }

    var o = {};
    o.a = "a";
    o.aa = "aa";
    o.aaa = "aaa";

    for (var n in o) {
      log(n);
      log(":");
      log(o[n]);
    }

    var F = function() {
      this.x = "x";
      this.y = "y";
      this.z = "z";
    };

    F.prototype.hello = function() {
      log("hello");
    };

    var f = new F();

    log(f.x);
    log(f.y);
    log(f.z);

    f.hello();
