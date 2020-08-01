export default function mumbleLetters(string = "") {
  if (!string) {
    return ""
  }

  let mumbledResult = [];

  for (let i = 0; i < string.length; i++) {
    let mumbledEntry = [];
    mumbledEntry.push(repeatLetter(string[i], i))
    mumbledResult.push(mumbledEntry.join(''));
  }
  return mumbledResult.join('-');
}

const repeatLetter = (letter = '', numberOfTimes = 0) => {
  let repeated = [];
  for (let i = 0; i < numberOfTimes + 1; i++) {
    i === 0 ? repeated.push(letter.toUpperCase()) : repeated.push(letter.toLowerCase());
  }

  return repeated.join('');
}