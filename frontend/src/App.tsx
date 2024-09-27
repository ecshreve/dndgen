import React from "react";
import CharacterSummary from "./CharacterSummary";

interface RaceDetails {
  id: string;
  indx: string;
  name: string;
  abilityBonuses: { bonus: number; abilityScore: { indx: string } }[];
  alignmentDesc: string;
  ageDesc: string;
  size: string;
  sizeDesc: string;
  speed: number;
  languages: { name: string }[];
  languageDesc: string;
  startingProficiencies: { indx: string }[];
  startingProficiencyOptions: {
    desc: string[];
    choose: number;
    proficiencies: { indx: string }[];
  };
}

interface ClassDetails {
  id: string;
  indx: string;
  name: string;
  savingThrows: { indx: string }[];
  proficiencies: { indx: string }[];
  proficiencyOptions: { desc: string }[];
  startingEquipment: { quantity: number; equipment: { indx: string } }[];
}

const App: React.FC = () => {
  return <CharacterSummary />;
};

export default App;
