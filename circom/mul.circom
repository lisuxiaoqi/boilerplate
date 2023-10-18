pragma circom 2.0.0;

template Mul() {
    signal input a;
    signal input b;
    signal output c;
    c <== a*b;

    log(123);
    log("bo");
 }

 component main = Mul();