const fs = require("fs");

const textIn = fs.readFileSync("./input.txt", "utf8");
const arr = textIn.split("\n");

const moveNumber = [];
const moveFrom = [];
const moveTo = [];

for (let k = 10; k < arr.length; k++) {
  moveNumber.push(Number(arr[k].slice(5, 7)));
  moveFrom.push(Number(arr[k].slice(12, 14)));
  moveTo.push(Number(arr[k].slice(17, 19)));
}

const readCrates = (total) => {
  for (let i = 1; i < 35; i += 4) {
    let crates = [];
    for (let j = 0; j < 8; j++) {
      if (arr[j][i] != " ") {
        crates.push(arr[j][i]);
      }
    }
    total.push(crates);
  }
};

// Part 1
const allCrates = [];
readCrates(allCrates);
for (let index = 0; index < moveNumber.length; index++) {
  for (let item = 0; item < moveNumber[index]; item++) {
    let temp = "";
    temp = allCrates[moveFrom[index] - 1].shift();
    allCrates[moveTo[index] - 1].unshift(temp);
  }
}

// Part 2
const allCrates2 = [];
readCrates(allCrates2);
for (let index = 0; index < moveNumber.length - 1; index++) {
  let temp = [];
  temp = allCrates2[moveFrom[index] - 1].slice(0, moveNumber[index]);
  allCrates2[moveFrom[index] - 1] = allCrates2[moveFrom[index] - 1].slice(
    moveNumber[index]
  );
  for (let j = temp.length - 1; j >= 0; j--) {
    allCrates2[moveTo[index] - 1].unshift(temp[j]);
  }
}

console.log("Part 1   Part 2 \n---------------");
for (let i = 0; i < allCrates.length; i++) {
  console.log(`${allCrates[i][0]}      |      ${allCrates2[i][0]}`);
}
