export const abilityScoreToModifier = (score: number): number => {
  if (score < 1) {
    return -5;
  }

  if (score > 30) {
    return 10;
  }

  const mod = Math.floor((score - 10) / 2);

  return mod;
};

export const calculateProfBonus = (level: number): number => {
  if (level < 5) {
    return 2;
  }

  if (level < 9) {
    return 3;
  }

  if (level < 13) {
    return 4;
  }

  if (level < 17) {
    return 5;
  }

  return 6;
};
