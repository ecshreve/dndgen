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
