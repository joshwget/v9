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
