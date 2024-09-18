// types.ts

export interface AttributeScores {
    strength: number;
    dexterity: number;
    constitution: number;
    intelligence: number;
    wisdom: number;
    charisma: number;
  }
  
  export interface CharacterFormValues {
    name: string;
    classId: string;
    raceId: string;
    alignment: string;
    attributes: AttributeScores;
  }
  
  export interface Option {
    id: string;
    name: string;
  }
  