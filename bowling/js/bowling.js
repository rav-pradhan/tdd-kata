const bowlingScore = (frames = "") => {
  let finalScore = 0;
  const sanitisedFrames = buildFramesAsNumberedScores(frames);
  for (let i = 0; i < sanitisedFrames.length; i++) {
    let frameScore = calculateFrameScore(sanitisedFrames, i);
    finalScore += frameScore;
  }
  return finalScore;
};

const buildFramesAsNumberedScores = (frames) => {
  return frames.split(" ").map((frame) => {
    let rollScores = [];
    for (let i = 0; i < frame.length; i++) {
      if (frame[i] == "/") {
        rollScores.push(10 - parseInt(frame[i - 1]));
      } else if (frame[i] == "X") {
        rollScores.push(10);
      } else {
        rollScores.push(parseInt(frame[i]));
      }
    }
    return rollScores;
  });
};

const calculateFrameScore = (frames, index) => {
  if (frameHasStrike(frames[index])) {
    return calculateStrikeBonusPoints(index, frames);
  }
  if (frameIsSpare(frames[index])) {
    return calculateSpareBonusPoints(index, frames);
  }
  return calculatePinsFallenInFrame(frames[index]);
};

const calculatePinsFallenInFrame = (frame) => {
  let pinsFallen = 0;
  frame.forEach((rollScore) => {
    pinsFallen += rollScore;
  });
  return pinsFallen;
};

const calculateSpareBonusPoints = (index, framesAsArray) => {
  if (isFinalFrame(index)) {
    return 10 + framesAsArray[index][2]
  } else {
    return 10 + framesAsArray[index + 1][0];
  }
};

const calculateStrikeBonusPoints = (index, framesAsArray) => {
  if (isFinalFrame(index)) {
    return 10 + framesAsArray[index][1] + framesAsArray[index][2]
  } else if (nextFrameHasStrike(index, framesAsArray)) {
    return 20 + framesAsArray[index+2][0]
  } else {
    return 10 + framesAsArray[index+1][0] + framesAsArray[index+1][1]
  }
};

const frameIsSpare = (frame) => {
  return frame[0] + frame[1] == 10;
};

const frameHasStrike = (frame) => {
  return frame[0] == 10;
};

const isFinalFrame = (index) => {
  return index == 9
}

const nextFrameHasStrike = (index, frames) => {
  return index < frames.length - 2 && frameHasStrike(frames[index+1])
}

export default bowlingScore;
