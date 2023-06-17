const fs = require("fs");

data = fs.readFileSync("./test.txt", "utf-8").split("\n");

// Part 1

// tailArray = [{ row: 0, col: 0 }];
// headArray = [{ row: 0, col: 0 }];

// const moveTail = () => {
//   if (
//     (Math.abs(headArray[headArray.length - 1].row - prevTailPos.row) > 1 &&
//       Math.abs(headArray[headArray.length - 1].col - prevTailPos.col) > 1) ||
//     Math.abs(headArray[headArray.length - 1].row - prevTailPos.row) > 1 ||
//     Math.abs(headArray[headArray.length - 1].col - prevTailPos.col) > 1
//   ) {
//     tailArray.push({
//       row: headArray[headArray.length - 2].row,
//       col: headArray[headArray.length - 2].col,
//     });
//   }
// };

// for (let i = 0; i < data.length; i++) {
//   let shift = 0;

//   // Get head shift value
//   if (data[i].length == 3) {
//     shift = Number(data[i][2]);
//   } else if (data[i].length == 4) {
//     shift = Number(data[i][2] + data[i][3]);
//   }
//   // Move head

//   for (let j = 0; j < shift; j++) {
//     prevHeadPos = headArray[headArray.length - 1];
//     prevTailPos = tailArray[tailArray.length - 1];

//     if (data[i][0] == "U") {
//       headArray.push({ row: prevHeadPos.row + 1, col: prevHeadPos.col });
//     } else if (data[i][0] == "D") {
//       headArray.push({ row: prevHeadPos.row - 1, col: prevHeadPos.col });
//     } else if (data[i][0] == "L") {
//       headArray.push({ row: prevHeadPos.row, col: prevHeadPos.col - 1 });
//     } else if (data[i][0] == "R") {
//       headArray.push({ row: prevHeadPos.row, col: prevHeadPos.col + 1 });
//     }

//     // Move tails
//     moveTail();

//     console.log(headArray);
//     console.log(tailArray);
//     const unique = [
//       ...new Set(
//         tailArray.map((item) => item.row.toString() + "_" + item.col.toString())
//       ),
//     ];
//     console.log(unique, unique.length);
//   }
// }

// Part 2

headArray = [{ row: 0, col: 0 }];
tailArray = [{ row: 0, col: 0 }];
arr = [
  { row: 0, col: 0 },
  { row: 0, col: 0 },
  { row: 0, col: 0 },
  { row: 0, col: 0 },
  { row: 0, col: 0 },
  { row: 0, col: 0 },
  { row: 0, col: 0 },
  { row: 0, col: 0 },
  { row: 0, col: 0 },
];
// arr3 = { row: 0, col: 0 };
// arr2 = { row: 0, col: 0 };
// arr4 = { row: 0, col: 0 };
// arr5 = { row: 0, col: 0 };
// arr6 = { row: 0, col: 0 };
// arr7 = { row: 0, col: 0 };
// arr8 = { row: 0, col: 0 };

const moveTail = (head, tail, move) => {
  if (
    (Math.abs(head.row - tail.row) > 1 && Math.abs(head.col - tail.col) > 1) ||
    Math.abs(head.row - tail.row) > 1 ||
    Math.abs(head.col - tail.col) > 1
  ) {
    tail.row = move.row;
    tail.col = move.col;
  }
  return tail;
};

for (let i = 0; i < data.length; i++) {
  let shift = 0;

  // Get head shift value
  if (data[i].length == 3) {
    shift = Number(data[i][2]);
  } else if (data[i].length == 4) {
    shift = Number(data[i][2] + data[i][3]);
  }
  // Move head

  for (let j = 0; j < shift; j++) {
    prevHeadPos = headArray[headArray.length - 1];
    let tempHead = prevHeadPos;

    if (data[i][0] == "U") {
      headArray.push({ row: prevHeadPos.row + 1, col: prevHeadPos.col });
    } else if (data[i][0] == "D") {
      headArray.push({ row: prevHeadPos.row - 1, col: prevHeadPos.col });
    } else if (data[i][0] == "L") {
      headArray.push({ row: prevHeadPos.row, col: prevHeadPos.col - 1 });
    } else if (data[i][0] == "R") {
      headArray.push({ row: prevHeadPos.row, col: prevHeadPos.col + 1 });
    }

    for (let k = 0; k < 9; k++) {
      if (k == 0) {
        tempHead = moveTail(
          headArray[headArray.length - 1],
          arr[k],
          headArray[headArray.length - 2]
        );
      } else if (k == 8) {
        tempHead = moveTail(arr[k - 1], arr[k], tempHead);
        tailArray.push(tempHead);
      } else {
        tempHead = moveTail(arr[k - 1], arr[k], tempHead);
      }
    }
  }
}
console.log(tailArray);
// if (
//   (Math.abs(
//     headArray[headArray.length - 1].row - arr1[arr1.length - 1].row
//   ) > 1 &&
//     Math.abs(
//       headArray[headArray.length - 1].col - arr1[arr1.length - 1].col
//     ) > 1) ||
//   Math.abs(
//     headArray[headArray.length - 1].row - arr1[arr1.length - 1].row
//   ) > 1 ||
//   Math.abs(
//     headArray[headArray.length - 1].col - arr1[arr1.length - 1].col
//   ) > 1
// ) {
//   arr1.push({
//     row: headArray[headArray.length - 2].row,
//     col: headArray[headArray.length - 2].col,
//   });
// }
// if (
//   (Math.abs(arr1[arr1.length - 1].row - arr2[arr2.length - 1].row) > 1 &&
//     Math.abs(arr1[arr1.length - 1].col - arr2[arr2.length - 1].col) > 1) ||
//   Math.abs(arr1[arr1.length - 1].row - arr2[arr2.length - 1].row) > 1 ||
//   Math.abs(arr1[arr1.length - 1].col - arr2[arr2.length - 1].col) > 1
// ) {
//   arr2.push({
//     row: arr1[arr1.length - 2].row,
//     col: arr1[arr1.length - 2].col,
//   });
// }
// if (
//   (Math.abs(arr2[arr2.length - 1].row - arr3[arr3.length - 1].row) > 1 &&
//     Math.abs(arr2[arr2.length - 1].col - arr3[arr3.length - 1].col) > 1) ||
//   Math.abs(arr2[arr2.length - 1].row - arr3[arr3.length - 1].row) > 1 ||
//   Math.abs(arr2[arr2.length - 1].col - arr3[arr3.length - 1].col) > 1
// ) {
//   arr3.push({
//     row: arr2[arr2.length - 2].row,
//     col: arr2[arr2.length - 2].col,
//   });
// }
// if (
//   (Math.abs(arr3[arr3.length - 1].row - arr4[arr4.length - 1].row) > 1 &&
//     Math.abs(arr3[arr3.length - 1].col - arr4[arr4.length - 1].col) > 1) ||
//   Math.abs(arr3[arr3.length - 1].row - arr4[arr4.length - 1].row) > 1 ||
//   Math.abs(arr3[arr3.length - 1].col - arr4[arr4.length - 1].col) > 1
// ) {
//   arr4.push({
//     row: arr3[arr3.length - 2].row,
//     col: arr3[arr3.length - 2].col,
//   });
// }
// if (
//   (Math.abs(arr4[arr4.length - 1].row - arr5[arr5.length - 1].row) > 1 &&
//     Math.abs(arr4[arr4.length - 1].col - arr5[arr5.length - 1].col) > 1) ||
//   Math.abs(arr4[arr4.length - 1].row - arr5[arr5.length - 1].row) > 1 ||
//   Math.abs(arr4[arr4.length - 1].col - arr5[arr5.length - 1].col) > 1
// ) {
//   arr5.push({
//     row: arr4[arr4.length - 2].row,
//     col: arr4[arr4.length - 2].col,
//   });
// }
// if (
//   (Math.abs(arr5[arr5.length - 1].row - arr6[arr6.length - 1].row) > 1 &&
//     Math.abs(arr5[arr5.length - 1].col - arr6[arr6.length - 1].col) > 1) ||
//   Math.abs(arr5[arr5.length - 1].row - arr6[arr6.length - 1].row) > 1 ||
//   Math.abs(arr5[arr5.length - 1].col - arr6[arr6.length - 1].col) > 1
// ) {
//   arr6.push({
//     row: arr5[arr5.length - 2].row,
//     col: arr5[arr5.length - 2].col,
//   });
// }
// if (
//   (Math.abs(arr6[arr6.length - 1].row - arr7[arr7.length - 1].row) > 1 &&
//     Math.abs(arr6[arr6.length - 1].col - arr7[arr7.length - 1].col) > 1) ||
//   Math.abs(arr6[arr6.length - 1].row - arr7[arr7.length - 1].row) > 1 ||
//   Math.abs(arr6[arr6.length - 1].col - arr7[arr7.length - 1].col) > 1
// ) {
//   arr7.push({
//     row: arr6[arr6.length - 2].row,
//     col: arr6[arr6.length - 2].col,
//   });
// }
// if (
//   (Math.abs(arr7[arr7.length - 1].row - arr8[arr8.length - 1].row) > 1 &&
//     Math.abs(arr7[arr7.length - 1].col - arr8[arr8.length - 1].col) > 1) ||
//   Math.abs(arr7[arr7.length - 1].row - arr8[arr8.length - 1].row) > 1 ||
//   Math.abs(arr7[arr7.length - 1].col - arr8[arr8.length - 1].col) > 1
// ) {
//   arr8.push({
//     row: arr7[arr7.length - 2].row,
//     col: arr7[arr7.length - 2].col,
//   });
// }
// if (
//   (Math.abs(
//     arr8[arr8.length - 1].row - tailArray[tailArray.length - 1].row
//   ) > 1 &&
//     Math.abs(
//       arr8[arr8.length - 1].col - tailArray[tailArray.length - 1].col
//     ) > 1) ||
//   Math.abs(
//     arr8[arr8.length - 1].row - tailArray[tailArray.length - 1].row
//   ) > 1 ||
//   Math.abs(
//     arr8[arr8.length - 1].col - tailArray[tailArray.length - 1].col
//   ) > 1
// ) {
//   tailArray.push({
//     row: arr8[arr8.length - 2].row,
//     col: arr8[arr8.length - 2].col,
//   });
// }
