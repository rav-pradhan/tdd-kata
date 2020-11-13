export default function mumbleLetters(string = "") {
  return buildMumbledLettersArray(string).join('-')
}

const buildMumbledLettersArray = (string) => {
  return string.split("").map((character, index) => {
    let mumbledEntry = []
    mumbledEntry.push(`${character.toUpperCase()}${character.toLowerCase().repeat(index)}`)
    return mumbledEntry
  })
}