/**
 * @author Miles Aube
 * @version 1.0.0
 * @date 2025-12-07
 * @fileoverview This program calculates base exponent using a While loop.
 */

// variables
let base: number = 0;
let exponent: number = 0;
let result: number = 1;
let counter: number = 0;

// input from the user
base = Number(prompt("Enter the base: ") || "0");
exponent = Number(prompt("Enter the exponent: ") || "0");

// calculate using a while loop
while (counter < exponent) {
  result = result * base;
  counter = counter + 1;
}

// output the answer
console.log("\n" + base + " raised to the power of " + exponent + " is: " + result);