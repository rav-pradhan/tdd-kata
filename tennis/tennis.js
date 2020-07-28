export default class TennisMatch {
  constructor(p1Name = "", p2Name = "") {
    this.playerOne = {
      name: p1Name,
      score: 0,
    };
    this.playerTwo = {
      name: p2Name,
      score: 0,
    };
    this.gameWinner = ""
  }

  playerOneWinsAPoint() {
    this.playerOne.score++;
  }

  playerTwoWinsAPoint() {
    this.playerTwo.score++;
  }

  getScore() {
    if (this.gameHasAWinner()) {
      return `Game won by ${this.gameWinner}`
    }

    if (this.gameIsInDeuceSituation()) {
      return this.calculateGamePointScenario();
    }

    return `${this.translateScore(this.playerOne.score)}-${this.translateScore(
      this.playerTwo.score
    )}`;
  }

  gameHasAWinner() {
    if (this.playerOne.score === 4 && this.playerTwo.score < 3) {
      this.gameWinner = this.playerOne.name
      return true
    }

    if (this.playerTwo.score === 4 && this.playerOne.score < 3) {
      this.gameWinner = this.playerTwo.name
      return true
    }
  }

  gameIsInDeuceSituation() {
    return this.playerOne.score >= 3 && this.playerTwo.score >= 3
  }

  calculateGamePointScenario() {
    if (this.playerOne.score === 4 && this.playerTwo.score === 3) {
      return "A-40";
    }

    if (this.playerOne.score === 3 && this.playerTwo.score === 4) {
      return "40-A";
    }

    if (this.playerOne.score === 4 && this.playerTwo.score === 4) {
      this.setScoresToDeuce()
      return "40-40"
    }
  }

  setScoresToDeuce() {
    this.playerOne.score = 3;
    this.playerTwo.score = 3;
  }

  translateScore(score = 0) {
    switch (score) {
      case 0:
        return "0";
      case 1:
        return "15";
      case 2:
        return "30";
      case 3:
        return "40";
      default:
        throw new Error("Score is out of bounds")
    }
  }
}
