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

export type Character = Node & {
  __typename?: 'Character';
  alignment?: Maybe<Scalars['String']['output']>;
  class: Class;
  id: Scalars['ID']['output'];
  level: Scalars['Int']['output'];
  name: Scalars['String']['output'];
  race: Race;
};

/** Ordering options for Character connections */
export type CharacterOrder = {
  /** The ordering direction. */
  direction?: OrderDirection;
  /** The field by which to order Characters. */
  field: CharacterOrderField;
};

/** Properties by which Character connections can be ordered. */
export enum CharacterOrderField {
  Name = 'NAME'
}

/**
 * CharacterWhereInput is used for filtering Character objects.
 * Input was generated by ent.
 */
export type CharacterWhereInput = {
  /** alignment field predicates */
  alignment?: InputMaybe<Scalars['String']['input']>;
  alignmentContains?: InputMaybe<Scalars['String']['input']>;
  alignmentContainsFold?: InputMaybe<Scalars['String']['input']>;
  alignmentEqualFold?: InputMaybe<Scalars['String']['input']>;
  alignmentGT?: InputMaybe<Scalars['String']['input']>;
  alignmentGTE?: InputMaybe<Scalars['String']['input']>;
  alignmentHasPrefix?: InputMaybe<Scalars['String']['input']>;
  alignmentHasSuffix?: InputMaybe<Scalars['String']['input']>;
  alignmentIn?: InputMaybe<Array<Scalars['String']['input']>>;
  alignmentIsNil?: InputMaybe<Scalars['Boolean']['input']>;
  alignmentLT?: InputMaybe<Scalars['String']['input']>;
  alignmentLTE?: InputMaybe<Scalars['String']['input']>;
  alignmentNEQ?: InputMaybe<Scalars['String']['input']>;
  alignmentNotIn?: InputMaybe<Array<Scalars['String']['input']>>;
  alignmentNotNil?: InputMaybe<Scalars['Boolean']['input']>;
  and?: InputMaybe<Array<CharacterWhereInput>>;
  /** class edge predicates */
  hasClass?: InputMaybe<Scalars['Boolean']['input']>;
  hasClassWith?: InputMaybe<Array<ClassWhereInput>>;
  /** race edge predicates */
  hasRace?: InputMaybe<Scalars['Boolean']['input']>;
  hasRaceWith?: InputMaybe<Array<RaceWhereInput>>;
  /** id field predicates */
  id?: InputMaybe<Scalars['ID']['input']>;
  idGT?: InputMaybe<Scalars['ID']['input']>;
  idGTE?: InputMaybe<Scalars['ID']['input']>;
  idIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  idLT?: InputMaybe<Scalars['ID']['input']>;
  idLTE?: InputMaybe<Scalars['ID']['input']>;
  idNEQ?: InputMaybe<Scalars['ID']['input']>;
  idNotIn?: InputMaybe<Array<Scalars['ID']['input']>>;
  /** level field predicates */
  level?: InputMaybe<Scalars['Int']['input']>;
  levelGT?: InputMaybe<Scalars['Int']['input']>;
  levelGTE?: InputMaybe<Scalars['Int']['input']>;
  levelIn?: InputMaybe<Array<Scalars['Int']['input']>>;
  levelLT?: InputMaybe<Scalars['Int']['input']>;
  levelLTE?: InputMaybe<Scalars['Int']['input']>;
  levelNEQ?: InputMaybe<Scalars['Int']['input']>;
  levelNotIn?: InputMaybe<Array<Scalars['Int']['input']>>;
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
  characters?: Maybe<Array<Character>>;
  id: Scalars['ID']['output'];
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
  Name = 'NAME'
}

/**
 * ClassWhereInput is used for filtering Class objects.
 * Input was generated by ent.
 */
export type ClassWhereInput = {
  and?: InputMaybe<Array<ClassWhereInput>>;
  /** characters edge predicates */
  hasCharacters?: InputMaybe<Scalars['Boolean']['input']>;
  hasCharactersWith?: InputMaybe<Array<CharacterWhereInput>>;
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
  not?: InputMaybe<ClassWhereInput>;
  or?: InputMaybe<Array<ClassWhereInput>>;
};

/**
 * CreateCharacterInput is used for create Character object.
 * Input was generated by ent.
 */
export type CreateCharacterInput = {
  alignment?: InputMaybe<Scalars['String']['input']>;
  classID: Scalars['ID']['input'];
  level?: InputMaybe<Scalars['Int']['input']>;
  name: Scalars['String']['input'];
  raceID: Scalars['ID']['input'];
};

/**
 * CreateClassInput is used for create Class object.
 * Input was generated by ent.
 */
export type CreateClassInput = {
  characterIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  name: Scalars['String']['input'];
};

/**
 * CreateRaceInput is used for create Race object.
 * Input was generated by ent.
 */
export type CreateRaceInput = {
  characterIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  name: Scalars['String']['input'];
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
  characters: Array<Character>;
  classes: Array<Class>;
  /** Fetches an object given its ID. */
  node?: Maybe<Node>;
  /** Lookup nodes by a list of IDs. */
  nodes: Array<Maybe<Node>>;
  races: Array<Race>;
};


export type QueryNodeArgs = {
  id: Scalars['ID']['input'];
};


export type QueryNodesArgs = {
  ids: Array<Scalars['ID']['input']>;
};

export type Race = Node & {
  __typename?: 'Race';
  characters?: Maybe<Array<Character>>;
  id: Scalars['ID']['output'];
  name: Scalars['String']['output'];
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
  Name = 'NAME'
}

/**
 * RaceWhereInput is used for filtering Race objects.
 * Input was generated by ent.
 */
export type RaceWhereInput = {
  and?: InputMaybe<Array<RaceWhereInput>>;
  /** characters edge predicates */
  hasCharacters?: InputMaybe<Scalars['Boolean']['input']>;
  hasCharactersWith?: InputMaybe<Array<CharacterWhereInput>>;
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
  not?: InputMaybe<RaceWhereInput>;
  or?: InputMaybe<Array<RaceWhereInput>>;
};

/**
 * UpdateCharacterInput is used for update Character object.
 * Input was generated by ent.
 */
export type UpdateCharacterInput = {
  alignment?: InputMaybe<Scalars['String']['input']>;
  classID?: InputMaybe<Scalars['ID']['input']>;
  clearAlignment?: InputMaybe<Scalars['Boolean']['input']>;
  level?: InputMaybe<Scalars['Int']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  raceID?: InputMaybe<Scalars['ID']['input']>;
};

/**
 * UpdateClassInput is used for update Class object.
 * Input was generated by ent.
 */
export type UpdateClassInput = {
  addCharacterIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  clearCharacters?: InputMaybe<Scalars['Boolean']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  removeCharacterIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
};

/**
 * UpdateRaceInput is used for update Race object.
 * Input was generated by ent.
 */
export type UpdateRaceInput = {
  addCharacterIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
  clearCharacters?: InputMaybe<Scalars['Boolean']['input']>;
  name?: InputMaybe<Scalars['String']['input']>;
  removeCharacterIDs?: InputMaybe<Array<Scalars['ID']['input']>>;
};

export type GetClassesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetClassesQuery = { __typename?: 'Query', classes: Array<{ __typename?: 'Class', id: string, name: string }> };


export const GetClassesDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"query","name":{"kind":"Name","value":"GetClasses"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"classes"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"id"}},{"kind":"Field","name":{"kind":"Name","value":"name"}}]}}]}}]} as unknown as DocumentNode<GetClassesQuery, GetClassesQueryVariables>;