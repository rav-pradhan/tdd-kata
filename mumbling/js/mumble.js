export default function mumbleLetters(string = "") {
  return buildMumbledLettersArray(string).join('-')
}

const buildMumbledLettersArray = (string = "") => {
  return string.split("").map((character, index) => {
    let mumbledEntry = [];
    mumbledEntry.push(repeatLetter(character, index))
    return mumbledEntry
  })
}

const repeatLetter = (letter = '', numberOfTimes = 0) => {
  let repeated = [];
  
  for (let i = 0; i < numberOfTimes + 1; i++) {
    i === 0 ? repeated.push(letter.toUpperCase()) : repeated.push(letter.toLowerCase());
  }

  return repeated.join('');
}