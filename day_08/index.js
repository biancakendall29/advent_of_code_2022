const fs = require("fs");

const textIn = fs.readFileSync("./input.txt", "utf8");
const arr = textIn.split("\n");

const getGrid = () => {
  for (let i = 0; i < arr.length; i++) {
    let row = [];
    for (let j = 0; j < arr[i].length; j++) {
      row.push(Number(arr[i][j]));
    }
    grid.push(row);
  }
};

const grid = [];
getGrid();
const length = grid[0].length;

// Part 1

const checkRight = (row, col) => {
  for (let i = col + 1; i < length; i++) {
    if (grid[row][i] >= grid[row][col]) {
      break;
    } else if (i == length - 1) {
      allVis.set(
        row.toString() + col.toString(),
        allVis.get(row.toString() + col.toString()) + 1 || 1
      );
    }
  }
};

const checkLeft = (row, col) => {
  for (let i = col - 1; i > 0; i--) {
    if (grid[row][i] >= grid[row][col]) {
      break;
    } else if (i == 1) {
      allVis.set(
        row.toString() + col.toString(),
        allVis.get(row.toString() + col.toString()) + 1 || 1
      );
    }
  }
};

const checkDown = (row, col) => {
  for (let i = row + 1; i < length; i++) {
    if (grid[i][col] >= grid[row][col]) {
      break;
    } else if (i == length - 1) {
      allVis.set(
        row.toString() + col.toString(),
        allVis.get(row.toString() + col.toString()) + 1 || 1
      );
    }
  }
};

const checkUp = (row, col) => {
  for (let i = row - 1; i > 0; i--) {
    if (grid[i][col] >= grid[row][col]) {
      break;
    } else if (i == 1) {
      allVis.set(
        row.toString() + col.toString(),
        allVis.get(row.toString() + col.toString()) + 1 || 1
      );
    }
  }
};

const loopGrid = () => {
  for (let row = 1; row < length; row++) {
    for (let col = 1; col < length; col++) {
      checkRight(row, col);
      checkLeft(row, col);
      checkDown(row, col);
      checkUp(row, col);
    }
  }
};

const allVis = new Map();
loopGrid();
console.log(allVis.size + 3 * length + (length - 1));

// Part 2

const checkRightScore = (row, col, counter) => {
  for (let i = col + 1; i < length; i++) {
    counter++;
    if (grid[row][i] >= grid[row][col]) {
      break;
    }
  }
  return counter;
};

const checkLeftScore = (row, col, counter) => {
  for (let i = col - 1; i >= 0; i--) {
    counter++;
    if (grid[row][i] >= grid[row][col]) {
      break;
    }
  }
  return counter;
};

const checkDownScore = (row, col, counter) => {
  for (let i = row + 1; i < length; i++) {
    counter++;
    if (grid[i][col] >= grid[row][col]) {
      break;
    }
  }
  return counter;
};

const checkUpScore = (row, col, counter) => {
  for (let i = row - 1; i >= 0; i--) {
    counter++;
    if (grid[i][col] >= grid[row][col]) {
      break;
    }
  }
  return counter;
};

const countScores = () => {
  for (let row = 1; row < length - 1; row++) {
    for (let col = 1; col < length - 1; col++) {
      let countRight = 0;
      let countLeft = 0;
      let countDown = 0;
      let countUp = 0;
      let key = row.toString() + col.toString();
      if (col < 10) {
        key = row.toString() + "0" + col.toString();
      }
      scenicScores.set(
        key,
        checkRightScore(row, col, countRight) *
          checkLeftScore(row, col, countLeft) *
          checkDownScore(row, col, countDown) *
          checkUpScore(row, col, countUp)
      );
    }
  }
};

const scenicScores = new Map();
countScores();
console.log(Math.max(...scenicScores.values()));
