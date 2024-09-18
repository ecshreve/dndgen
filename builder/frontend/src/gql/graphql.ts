/* eslint-disable */
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
  /**
   * Define a Relay Cursor type:
   * https://relay.dev/graphql/connections.htm#sec-Cursor
   */
  Cursor: { input: any; output: any; }
};

export type AbilityScore = Node & {
  __typename?: 'AbilityScore';
  abbr: AbilityScoreAbbr;
  desc: Array<Scalars['String']['output']>;
  id: Scalars['ID']['output'];
  indx: Scalars['String']['output'];
  name: Scalars['String']['output'];
};

/** AbilityScoreAbbr is enum for the field abbr */
export enum AbilityScoreAbbr {
  Cha = 'CHA',
  Con = 'CON',
  Dex = 'DEX',
  Int = 'INT',
  Str = 'STR',
  Wis = 'WIS'
}

/** Ordering options for AbilityScore connections */
export type AbilityScoreOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order AbilityScores. */
  field: AbilityScoreOrderField;
};

/** Properties by which AbilityScore connections can be ordered. */
export enum AbilityScoreOrderField {
  Indx = 'INDX',
  Name = 'NAME'
}

/**
 * AbilityScoreWhereInput is used for filtering AbilityScore objects.
 * Input was generated by ent.
 */
export type AbilityScoreWhereInput = {
  /** abbr field predicates */
  abbr?: InputMaybe<AbilityScoreAbbr>;
  abbrIn?: InputMaybe<Array<AbilityScoreAbbr>>;
  abbrNEQ?: InputMaybe<AbilityScoreAbbr>;
  abbrNotIn?: InputMaybe<Array<AbilityScoreAbbr>>;
  and?: InputMaybe<Array<AbilityScoreWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** indx field predicates */
  indx?: InputMaybe<Scalars['String']['input']>;
  indxContains?: InputMaybe<Scalars['String']['input']>;
  indxContainsFold?: InputMaybe<Scalars['String']['input']>;
  indxEqualFold?: InputMaybe<Scalars['String']['input']>;
  indxGT?: InputMaybe<Scalars['String']['input']>;
  indxGTE?: InputMaybe<Scalars['String']['input']>;
  indxHasPrefix?: InputMaybe<Scalars['String']['input']>;
  indxHasSuffix?: InputMaybe<Scalars['String']['input']>;
  indxIn?: InputMaybe<Array<Scalars['String']['input']>>;
  indxLT?: InputMaybe<Scalars['String']['input']>;
  indxLTE?: InputMaybe<Scalars['String']['input']>;
  indxNEQ?: InputMaybe<Scalars['String']['input']>;
  indxNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<AbilityScoreWhereInput>;
  or?: InputMaybe<Array<AbilityScoreWhereInput>>;
};

export type Alignment = Node & {
  __typename?: 'Alignment';
  abbr: Scalars['String']['output'];
  desc: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  indx: Scalars['String']['output'];
  name: Scalars['String']['output'];
};

/** Ordering options for Alignment connections */
export type AlignmentOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Alignments. */
  field: AlignmentOrderField;
};

/** Properties by which Alignment connections can be ordered. */
export enum AlignmentOrderField {
  Indx = 'INDX',
  Name = 'NAME'
}

/**
 * AlignmentWhereInput is used for filtering Alignment objects.
 * Input was generated by ent.
 */
export type AlignmentWhereInput = {
  /** abbr field predicates */
  abbr?: InputMaybe<Scalars['String']['input']>;
  abbrContains?: InputMaybe<Scalars['String']['input']>;
  abbrContainsFold?: InputMaybe<Scalars['String']['input']>;
  abbrEqualFold?: InputMaybe<Scalars['String']['input']>;
  abbrGT?: InputMaybe<Scalars['String']['input']>;
  abbrGTE?: InputMaybe<Scalars['String']['input']>;
  abbrHasPrefix?: InputMaybe<Scalars['String']['input']>;
  abbrHasSuffix?: InputMaybe<Scalars['String']['input']>;
  abbrIn?: InputMaybe<Array<Scalars['String']['input']>>;
  abbrLT?: InputMaybe<Scalars['String']['input']>;
  abbrLTE?: InputMaybe<Scalars['String']['input']>;
  abbrNEQ?: InputMaybe<Scalars['String']['input']>;
  abbrNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  and?: InputMaybe<Array<AlignmentWhereInput>>;
  /** desc field predicates */
  desc?: InputMaybe<Scalars['String']['input']>;
  descContains?: InputMaybe<Scalars['String']['input']>;
  descContainsFold?: InputMaybe<Scalars['String']['input']>;
  descEqualFold?: InputMaybe<Scalars['String']['input']>;
  descGT?: InputMaybe<Scalars['String']['input']>;
  descGTE?: InputMaybe<Scalars['String']['input']>;
  descHasPrefix?: InputMaybe<Scalars['String']['input']>;
  descHasSuffix?: InputMaybe<Scalars['String']['input']>;
  descIn?: InputMaybe<Array<Scalars['String']['input']>>;
  descLT?: InputMaybe<Scalars['String']['input']>;
  descLTE?: InputMaybe<Scalars['String']['input']>;
  descNEQ?: InputMaybe<Scalars['String']['input']>;
  descNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** indx field predicates */
  indx?: InputMaybe<Scalars['String']['input']>;
  indxContains?: InputMaybe<Scalars['String']['input']>;
  indxContainsFold?: InputMaybe<Scalars['String']['input']>;
  indxEqualFold?: InputMaybe<Scalars['String']['input']>;
  indxGT?: InputMaybe<Scalars['String']['input']>;
  indxGTE?: InputMaybe<Scalars['String']['input']>;
  indxHasPrefix?: InputMaybe<Scalars['String']['input']>;
  indxHasSuffix?: InputMaybe<Scalars['String']['input']>;
  indxIn?: InputMaybe<Array<Scalars['String']['input']>>;
  indxLT?: InputMaybe<Scalars['String']['input']>;
  indxLTE?: InputMaybe<Scalars['String']['input']>;
  indxNEQ?: InputMaybe<Scalars['String']['input']>;
  indxNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<AlignmentWhereInput>;
  or?: InputMaybe<Array<AlignmentWhereInput>>;
};

export type Character = Node & {
  __typename?: 'Character';
  age: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
};

/**
 * CharacterWhereInput is used for filtering Character objects.
 * Input was generated by ent.
 */
export type CharacterWhereInput = {
  /** age field predicates */
  age?: InputMaybe<Scalars['Int']['input']>;
  ageGT?: InputMaybe<Scalars['Int']['input']>;
  ageGTE?: InputMaybe<Scalars['Int']['input']>;
  ageIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  ageLT?: InputMaybe<Scalars['Int']['input']>;
  ageLTE?: InputMaybe<Scalars['Int']['input']>;
  ageNEQ?: InputMaybe<Scalars['Int']['input']>;
  ageNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  and?: InputMaybe<Array<CharacterWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<CharacterWhereInput>;
  or?: InputMaybe<Array<CharacterWhereInput>>;
};

export type Class = Node & {
  __typename?: 'Class';
  hitDie: Scalars['Int']['output'];
  id: Scalars['ID']['output'];
  indx: Scalars['String']['output'];
  name: Scalars['String']['output'];
};

/** Ordering options for Class connections */
export type ClassOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Classes. */
  field: ClassOrderField;
};

/** Properties by which Class connections can be ordered. */
export enum ClassOrderField {
  Indx = 'INDX',
  Name = 'NAME'
}

/**
 * ClassWhereInput is used for filtering Class objects.
 * Input was generated by ent.
 */
export type ClassWhereInput = {
  and?: InputMaybe<Array<ClassWhereInput>>;
  /** hit_die field predicates */
  hitDie?: InputMaybe<Scalars['Int']['input']>;
  hitDieGT?: InputMaybe<Scalars['Int']['input']>;
  hitDieGTE?: InputMaybe<Scalars['Int']['input']>;
  hitDieIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  hitDieLT?: InputMaybe<Scalars['Int']['input']>;
  hitDieLTE?: InputMaybe<Scalars['Int']['input']>;
  hitDieNEQ?: InputMaybe<Scalars['Int']['input']>;
  hitDieNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** indx field predicates */
  indx?: InputMaybe<Scalars['String']['input']>;
  indxContains?: InputMaybe<Scalars['String']['input']>;
  indxContainsFold?: InputMaybe<Scalars['String']['input']>;
  indxEqualFold?: InputMaybe<Scalars['String']['input']>;
  indxGT?: InputMaybe<Scalars['String']['input']>;
  indxGTE?: InputMaybe<Scalars['String']['input']>;
  indxHasPrefix?: InputMaybe<Scalars['String']['input']>;
  indxHasSuffix?: InputMaybe<Scalars['String']['input']>;
  indxIn?: InputMaybe<Array<Scalars['String']['input']>>;
  indxLT?: InputMaybe<Scalars['String']['input']>;
  indxLTE?: InputMaybe<Scalars['String']['input']>;
  indxNEQ?: InputMaybe<Scalars['String']['input']>;
  indxNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<ClassWhereInput>;
  or?: InputMaybe<Array<ClassWhereInput>>;
};

/**
 * CreateAbilityScoreInput is used for create AbilityScore object.
 * Input was generated by ent.
 */
export type CreateAbilityScoreInput = {
  abbr: AbilityScoreAbbr;
  desc: Array<Scalars['String']['input']>;
  indx: Scalars['String']['input'];
  name: Scalars['String']['input'];
};

/**
 * CreateAlignmentInput is used for create Alignment object.
 * Input was generated by ent.
 */
export type CreateAlignmentInput = {
  abbr: Scalars['String']['input'];
  desc: Scalars['String']['input'];
  indx: Scalars['String']['input'];
  name: Scalars['String']['input'];
};

/**
 * CreateClassInput is used for create Class object.
 * Input was generated by ent.
 */
export type CreateClassInput = {
  hitDie: Scalars['Int']['input'];
  indx: Scalars['String']['input'];
  name: Scalars['String']['input'];
};

/**
 * CreateLanguageInput is used for create Language object.
 * Input was generated by ent.
 */
export type CreateLanguageInput = {
  indx: Scalars['String']['input'];
  name: Scalars['String']['input'];
  script: LanguageScript;
  type: LanguageType;
};

/**
 * CreateMagicSchoolInput is used for create MagicSchool object.
 * Input was generated by ent.
 */
export type CreateMagicSchoolInput = {
  desc: Scalars['String']['input'];
  indx: Scalars['String']['input'];
  name: Scalars['String']['input'];
};

/**
 * CreateRaceInput is used for create Race object.
 * Input was generated by ent.
 */
export type CreateRaceInput = {
  age: Scalars['String']['input'];
  indx: Scalars['String']['input'];
  name: Scalars['String']['input'];
  size: RaceSize;
  sizeDescription: Scalars['String']['input'];
  speed: Scalars['Int']['input'];
};

/**
 * CreateSkillInput is used for create Skill object.
 * Input was generated by ent.
 */
export type CreateSkillInput = {
  indx: Scalars['String']['input'];
  name: Scalars['String']['input'];
};

export type Language = Node & {
  __typename?: 'Language';
  id: Scalars['ID']['output'];
  indx: Scalars['String']['output'];
  name: Scalars['String']['output'];
  script: LanguageScript;
  type: LanguageType;
};

/** Ordering options for Language connections */
export type LanguageOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Languages. */
  field: LanguageOrderField;
};

/** Properties by which Language connections can be ordered. */
export enum LanguageOrderField {
  Indx = 'INDX',
  Name = 'NAME'
}

/** LanguageScript is enum for the field script */
export enum LanguageScript {
  Abyssal = 'Abyssal',
  Celestial = 'Celestial',
  Common = 'Common',
  Draconic = 'Draconic',
  Drow = 'Drow',
  Dwarvish = 'Dwarvish',
  Elvish = 'Elvish',
  Giant = 'Giant',
  Gnomish = 'Gnomish',
  Goblin = 'Goblin',
  Halfling = 'Halfling',
  Infernal = 'Infernal',
  Language = 'Language',
  Orc = 'Orc',
  Primordial = 'Primordial',
  Sign = 'Sign',
  Sylvan = 'Sylvan',
  Undercommon = 'Undercommon'
}

/** LanguageType is enum for the field type */
export enum LanguageType {
  Exotic = 'EXOTIC',
  Standard = 'STANDARD'
}

/**
 * LanguageWhereInput is used for filtering Language objects.
 * Input was generated by ent.
 */
export type LanguageWhereInput = {
  and?: InputMaybe<Array<LanguageWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** indx field predicates */
  indx?: InputMaybe<Scalars['String']['input']>;
  indxContains?: InputMaybe<Scalars['String']['input']>;
  indxContainsFold?: InputMaybe<Scalars['String']['input']>;
  indxEqualFold?: InputMaybe<Scalars['String']['input']>;
  indxGT?: InputMaybe<Scalars['String']['input']>;
  indxGTE?: InputMaybe<Scalars['String']['input']>;
  indxHasPrefix?: InputMaybe<Scalars['String']['input']>;
  indxHasSuffix?: InputMaybe<Scalars['String']['input']>;
  indxIn?: InputMaybe<Array<Scalars['String']['input']>>;
  indxLT?: InputMaybe<Scalars['String']['input']>;
  indxLTE?: InputMaybe<Scalars['String']['input']>;
  indxNEQ?: InputMaybe<Scalars['String']['input']>;
  indxNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<LanguageWhereInput>;
  or?: InputMaybe<Array<LanguageWhereInput>>;
  /** script field predicates */
  script?: InputMaybe<LanguageScript>;
  scriptIn?: InputMaybe<Array<LanguageScript>>;
  scriptNEQ?: InputMaybe<LanguageScript>;
  scriptNotIn?: InputMaybe<Array<LanguageScript>>;
  /** type field predicates */
  type?: InputMaybe<LanguageType>;
  typeIn?: InputMaybe<Array<LanguageType>>;
  typeNEQ?: InputMaybe<LanguageType>;
  typeNotIn?: InputMaybe<Array<LanguageType>>;
};

export type MagicSchool = Node & {
  __typename?: 'MagicSchool';
  desc: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  indx: Scalars['String']['output'];
  name: Scalars['String']['output'];
};

/** Ordering options for MagicSchool connections */
export type MagicSchoolOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order MagicSchools. */
  field: MagicSchoolOrderField;
};

/** Properties by which MagicSchool connections can be ordered. */
export enum MagicSchoolOrderField {
  Indx = 'INDX',
  Name = 'NAME'
}

/**
 * MagicSchoolWhereInput is used for filtering MagicSchool objects.
 * Input was generated by ent.
 */
export type MagicSchoolWhereInput = {
  and?: InputMaybe<Array<MagicSchoolWhereInput>>;
  /** desc field predicates */
  desc?: InputMaybe<Scalars['String']['input']>;
  descContains?: InputMaybe<Scalars['String']['input']>;
  descContainsFold?: InputMaybe<Scalars['String']['input']>;
  descEqualFold?: InputMaybe<Scalars['String']['input']>;
  descGT?: InputMaybe<Scalars['String']['input']>;
  descGTE?: InputMaybe<Scalars['String']['input']>;
  descHasPrefix?: InputMaybe<Scalars['String']['input']>;
  descHasSuffix?: InputMaybe<Scalars['String']['input']>;
  descIn?: InputMaybe<Array<Scalars['String']['input']>>;
  descLT?: InputMaybe<Scalars['String']['input']>;
  descLTE?: InputMaybe<Scalars['String']['input']>;
  descNEQ?: InputMaybe<Scalars['String']['input']>;
  descNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** indx field predicates */
  indx?: InputMaybe<Scalars['String']['input']>;
  indxContains?: InputMaybe<Scalars['String']['input']>;
  indxContainsFold?: InputMaybe<Scalars['String']['input']>;
  indxEqualFold?: InputMaybe<Scalars['String']['input']>;
  indxGT?: InputMaybe<Scalars['String']['input']>;
  indxGTE?: InputMaybe<Scalars['String']['input']>;
  indxHasPrefix?: InputMaybe<Scalars['String']['input']>;
  indxHasSuffix?: InputMaybe<Scalars['String']['input']>;
  indxIn?: InputMaybe<Array<Scalars['String']['input']>>;
  indxLT?: InputMaybe<Scalars['String']['input']>;
  indxLTE?: InputMaybe<Scalars['String']['input']>;
  indxNEQ?: InputMaybe<Scalars['String']['input']>;
  indxNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<MagicSchoolWhereInput>;
  or?: InputMaybe<Array<MagicSchoolWhereInput>>;
};

/**
 * An object with an ID.
 * Follows the [Relay Global Object Identification Specification](https://relay.dev/graphql/objectidentification.htm)
 */
export type Node = {
  /** The id of the object. */
  id: Scalars['ID']['output'];
};

/** Possible directions in which to order a list of items when provided an `orderBy` argument. */
export enum OrderDirection {
  /** Specifies an ascending order for a given `orderBy` argument. */
  Asc = 'ASC',
  /** Specifies a descending order for a given `orderBy` argument. */
  Desc = 'DESC'
}

/**
 * Information about pagination in a connection.
 * https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
 */
export type PageInfo = {
  __typename?: 'PageInfo';
  /** When paginating forwards, the cursor to continue. */
  endCursor?: Maybe<Scalars['Cursor']['output']>;
  /** When paginating forwards, are there more items? */
  hasNextPage: Scalars['Boolean']['output'];
  /** When paginating backwards, are there more items? */
  hasPreviousPage: Scalars['Boolean']['output'];
  /** When paginating backwards, the cursor to continue. */
  startCursor?: Maybe<Scalars['Cursor']['output']>;
};

export type Query = {
  __typename?: 'Query';
  abilityScores: Array<AbilityScore>;
  alignments: Array<Alignment>;
  classes: Array<Class>;
  languages: Array<Language>;
  magicSchools: Array<MagicSchool>;
  /** Fetches an object given its ID. */
  node?: Maybe<Node>;
  /** Lookup nodes by a list of IDs. */
  nodes: Array<Maybe<Node>>;
  races: Array<Race>;
  skills: Array<Skill>;
};


export type QueryNodeArgs = {
  id: Scalars['ID']['input'];
};


export type QueryNodesArgs = {
  ids: Array<Scalars['ID']['input']>;
};

export type Race = Node & {
  __typename?: 'Race';
  age: Scalars['String']['output'];
  id: Scalars['ID']['output'];
  indx: Scalars['String']['output'];
  name: Scalars['String']['output'];
  size: RaceSize;
  sizeDescription: Scalars['String']['output'];
  speed: Scalars['Int']['output'];
};

/** Ordering options for Race connections */
export type RaceOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Races. */
  field: RaceOrderField;
};

/** Properties by which Race connections can be ordered. */
export enum RaceOrderField {
  Indx = 'INDX',
  Name = 'NAME'
}

/** RaceSize is enum for the field size */
export enum RaceSize {
  Gargantuan = 'Gargantuan',
  Huge = 'Huge',
  Large = 'Large',
  Medium = 'Medium',
  Small = 'Small',
  Tiny = 'Tiny'
}

/**
 * RaceWhereInput is used for filtering Race objects.
 * Input was generated by ent.
 */
export type RaceWhereInput = {
  /** age field predicates */
  age?: InputMaybe<Scalars['String']['input']>;
  ageContains?: InputMaybe<Scalars['String']['input']>;
  ageContainsFold?: InputMaybe<Scalars['String']['input']>;
  ageEqualFold?: InputMaybe<Scalars['String']['input']>;
  ageGT?: InputMaybe<Scalars['String']['input']>;
  ageGTE?: InputMaybe<Scalars['String']['input']>;
  ageHasPrefix?: InputMaybe<Scalars['String']['input']>;
  ageHasSuffix?: InputMaybe<Scalars['String']['input']>;
  ageIn?: InputMaybe<Array<Scalars['String']['input']>>;
  ageLT?: InputMaybe<Scalars['String']['input']>;
  ageLTE?: InputMaybe<Scalars['String']['input']>;
  ageNEQ?: InputMaybe<Scalars['String']['input']>;
  ageNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  and?: InputMaybe<Array<RaceWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** indx field predicates */
  indx?: InputMaybe<Scalars['String']['input']>;
  indxContains?: InputMaybe<Scalars['String']['input']>;
  indxContainsFold?: InputMaybe<Scalars['String']['input']>;
  indxEqualFold?: InputMaybe<Scalars['String']['input']>;
  indxGT?: InputMaybe<Scalars['String']['input']>;
  indxGTE?: InputMaybe<Scalars['String']['input']>;
  indxHasPrefix?: InputMaybe<Scalars['String']['input']>;
  indxHasSuffix?: InputMaybe<Scalars['String']['input']>;
  indxIn?: InputMaybe<Array<Scalars['String']['input']>>;
  indxLT?: InputMaybe<Scalars['String']['input']>;
  indxLTE?: InputMaybe<Scalars['String']['input']>;
  indxNEQ?: InputMaybe<Scalars['String']['input']>;
  indxNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<RaceWhereInput>;
  or?: InputMaybe<Array<RaceWhereInput>>;
  /** size field predicates */
  size?: InputMaybe<RaceSize>;
  /** size_description field predicates */
  sizeDescription?: InputMaybe<Scalars['String']['input']>;
  sizeDescriptionContains?: InputMaybe<Scalars['String']['input']>;
  sizeDescriptionContainsFold?: InputMaybe<Scalars['String']['input']>;
  sizeDescriptionEqualFold?: InputMaybe<Scalars['String']['input']>;
  sizeDescriptionGT?: InputMaybe<Scalars['String']['input']>;
  sizeDescriptionGTE?: InputMaybe<Scalars['String']['input']>;
  sizeDescriptionHasPrefix?: InputMaybe<Scalars['String']['input']>;
  sizeDescriptionHasSuffix?: InputMaybe<Scalars['String']['input']>;
  sizeDescriptionIn?: InputMaybe<Array<Scalars['String']['input']>>;
  sizeDescriptionLT?: InputMaybe<Scalars['String']['input']>;
  sizeDescriptionLTE?: InputMaybe<Scalars['String']['input']>;
  sizeDescriptionNEQ?: InputMaybe<Scalars['String']['input']>;
  sizeDescriptionNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  sizeIn?: InputMaybe<Array<RaceSize>>;
  sizeNEQ?: InputMaybe<RaceSize>;
  sizeNotIn?: InputMaybe<Array<RaceSize>>;
  /** speed field predicates */
  speed?: InputMaybe<Scalars['Int']['input']>;
  speedGT?: InputMaybe<Scalars['Int']['input']>;
  speedGTE?: InputMaybe<Scalars['Int']['input']>;
  speedIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  speedLT?: InputMaybe<Scalars['Int']['input']>;
  speedLTE?: InputMaybe<Scalars['Int']['input']>;
  speedNEQ?: InputMaybe<Scalars['Int']['input']>;
  speedNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
};

export type Skill = Node & {
  __typename?: 'Skill';
  id: Scalars['ID']['output'];
  indx: Scalars['String']['output'];
  name: Scalars['String']['output'];
};

/** Ordering options for Skill connections */
export type SkillOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Skills. */
  field: SkillOrderField;
};

/** Properties by which Skill connections can be ordered. */
export enum SkillOrderField {
  Indx = 'INDX',
  Name = 'NAME'
}

/**
 * SkillWhereInput is used for filtering Skill objects.
 * Input was generated by ent.
 */
export type SkillWhereInput = {
  and?: InputMaybe<Array<SkillWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** indx field predicates */
  indx?: InputMaybe<Scalars['String']['input']>;
  indxContains?: InputMaybe<Scalars['String']['input']>;
  indxContainsFold?: InputMaybe<Scalars['String']['input']>;
  indxEqualFold?: InputMaybe<Scalars['String']['input']>;
  indxGT?: InputMaybe<Scalars['String']['input']>;
  indxGTE?: InputMaybe<Scalars['String']['input']>;
  indxHasPrefix?: InputMaybe<Scalars['String']['input']>;
  indxHasSuffix?: InputMaybe<Scalars['String']['input']>;
  indxIn?: InputMaybe<Array<Scalars['String']['input']>>;
  indxLT?: InputMaybe<Scalars['String']['input']>;
  indxLTE?: InputMaybe<Scalars['String']['input']>;
  indxNEQ?: InputMaybe<Scalars['String']['input']>;
  indxNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  /** name field predicates */
  name?: InputMaybe<Scalars['String']['input']>;
  nameContains?: InputMaybe<Scalars['String']['input']>;
  nameContainsFold?: InputMaybe<Scalars['String']['input']>;
  nameEqualFold?: InputMaybe<Scalars['String']['input']>;
  nameGT?: InputMaybe<Scalars['String']['input']>;
  nameGTE?: InputMaybe<Scalars['String']['input']>;
  nameHasPrefix?: InputMaybe<Scalars['String']['input']>;
  nameHasSuffix?: InputMaybe<Scalars['String']['input']>;
  nameIn?: InputMaybe<Array<Scalars['String']['input']>>;
  nameLT?: InputMaybe<Scalars['String']['input']>;
  nameLTE?: InputMaybe<Scalars['String']['input']>;
  nameNEQ?: InputMaybe<Scalars['String']['input']>;
  nameNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  not?: InputMaybe<SkillWhereInput>;
  or?: InputMaybe<Array<SkillWhereInput>>;
};

/**
 * UpdateAbilityScoreInput is used for update AbilityScore object.
 * Input was generated by ent.
 */
export type UpdateAbilityScoreInput = {
  abbr?: InputMaybe<AbilityScoreAbbr>;
  appendDesc?: InputMaybe<Array<Scalars['String']['input']>>;
  desc?: InputMaybe<Array<Scalars['String']['input']>>;
  indx?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateAlignmentInput is used for update Alignment object.
 * Input was generated by ent.
 */
export type UpdateAlignmentInput = {
  abbr?: InputMaybe<Scalars['String']['input']>;
  desc?: InputMaybe<Scalars['String']['input']>;
  indx?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateClassInput is used for update Class object.
 * Input was generated by ent.
 */
export type UpdateClassInput = {
  hitDie?: InputMaybe<Scalars['Int']['input']>;
  indx?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateLanguageInput is used for update Language object.
 * Input was generated by ent.
 */
export type UpdateLanguageInput = {
  indx?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  script?: InputMaybe<LanguageScript>;
  type?: InputMaybe<LanguageType>;
};

/**
 * UpdateMagicSchoolInput is used for update MagicSchool object.
 * Input was generated by ent.
 */
export type UpdateMagicSchoolInput = {
  desc?: InputMaybe<Scalars['String']['input']>;
  indx?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
};

/**
 * UpdateRaceInput is used for update Race object.
 * Input was generated by ent.
 */
export type UpdateRaceInput = {
  age?: InputMaybe<Scalars['String']['input']>;
  indx?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  size?: InputMaybe<RaceSize>;
  sizeDescription?: InputMaybe<Scalars['String']['input']>;
  speed?: InputMaybe<Scalars['Int']['input']>;
};

/**
 * UpdateSkillInput is used for update Skill object.
 * Input was generated by ent.
 */
export type UpdateSkillInput = {
  indx?: InputMaybe<Scalars['String']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
};

export type GetClassesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetClassesQuery = { __typename?: 'Query', classes: Array<{ __typename?: 'Class', id: string, name: string }> };

export type GetRacesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetRacesQuery = { __typename?: 'Query', races: Array<{ __typename?: 'Race', id: string, name: string }> };

export type GetAlignmentsQueryVariables = Exact<{ [key: string]: never; }>;


export type GetAlignmentsQuery = { __typename?: 'Query', alignments: Array<{ __typename?: 'Alignment', id: string, name: string }> };

export type GetAbilityScoresQueryVariables = Exact<{ [key: string]: never; }>;


export type GetAbilityScoresQuery = { __typename?: 'Query', abilityScores: Array<{ __typename?: 'AbilityScore', id: string, name: string }> };


export const GetClassesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetClasses"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"classes"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<GetClassesQuery, GetClassesQueryVariables>;
export const GetRacesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetRaces"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"races"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<GetRacesQuery, GetRacesQueryVariables>;
export const GetAlignmentsDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetAlignments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"alignments"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<GetAlignmentsQuery, GetAlignmentsQueryVariables>;
export const GetAbilityScoresDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetAbilityScores"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"abilityScores"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<GetAbilityScoresQuery, GetAbilityScoresQueryVariables>;