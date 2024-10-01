/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n  fragment CharacterAbilityScoreFragment on CharacterAbilityScore {\n    id\n    score\n    modifier\n    abilityScore {\n      id\n      indx\n      name\n    }\n    characterSkills {\n      id\n      proficient\n      skill {\n        id\n        indx\n        name\n      }\n    }\n  }\n": types.CharacterAbilityScoreFragmentFragmentDoc,
    "\n  query GetCharacters {\n    characters {\n      edges {\n        node {\n          id\n          name\n          age\n          level\n          proficiencyBonus\n          class {\n            id\n            indx\n            name\n          }\n          race {\n            id\n            indx\n            name\n          }\n          alignment {\n            id\n            indx\n            name\n          }\n          characterAbilityScores {\n            id\n            score\n            modifier\n            abilityScore {\n              id\n              indx\n              name\n            }\n          }\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n": types.GetCharactersDocument,
    "\n  query GetClasses {\n    classes {\n      id\n      indx\n      name\n    }\n  }\n": types.GetClassesDocument,
    "\n  query GetRaces {\n    races {\n      id\n      indx\n      name\n    }\n  }\n": types.GetRacesDocument,
    "\n  mutation UpdateCharacter($updateCharacterId: ID!, $input: UpdateCharacterInput!) {\n    updateCharacter(id: $updateCharacterId, input: $input) {\n      id\n      name\n      age\n      level\n      race {\n        id\n        indx\n        name\n      }\n      class {\n        id\n        indx\n        name\n      }\n    }\n  }\n": types.UpdateCharacterDocument,
};

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = gql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function gql(source: string): unknown;

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  fragment CharacterAbilityScoreFragment on CharacterAbilityScore {\n    id\n    score\n    modifier\n    abilityScore {\n      id\n      indx\n      name\n    }\n    characterSkills {\n      id\n      proficient\n      skill {\n        id\n        indx\n        name\n      }\n    }\n  }\n"): (typeof documents)["\n  fragment CharacterAbilityScoreFragment on CharacterAbilityScore {\n    id\n    score\n    modifier\n    abilityScore {\n      id\n      indx\n      name\n    }\n    characterSkills {\n      id\n      proficient\n      skill {\n        id\n        indx\n        name\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetCharacters {\n    characters {\n      edges {\n        node {\n          id\n          name\n          age\n          level\n          proficiencyBonus\n          class {\n            id\n            indx\n            name\n          }\n          race {\n            id\n            indx\n            name\n          }\n          alignment {\n            id\n            indx\n            name\n          }\n          characterAbilityScores {\n            id\n            score\n            modifier\n            abilityScore {\n              id\n              indx\n              name\n            }\n          }\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query GetCharacters {\n    characters {\n      edges {\n        node {\n          id\n          name\n          age\n          level\n          proficiencyBonus\n          class {\n            id\n            indx\n            name\n          }\n          race {\n            id\n            indx\n            name\n          }\n          alignment {\n            id\n            indx\n            name\n          }\n          characterAbilityScores {\n            id\n            score\n            modifier\n            abilityScore {\n              id\n              indx\n              name\n            }\n          }\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetClasses {\n    classes {\n      id\n      indx\n      name\n    }\n  }\n"): (typeof documents)["\n  query GetClasses {\n    classes {\n      id\n      indx\n      name\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetRaces {\n    races {\n      id\n      indx\n      name\n    }\n  }\n"): (typeof documents)["\n  query GetRaces {\n    races {\n      id\n      indx\n      name\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation UpdateCharacter($updateCharacterId: ID!, $input: UpdateCharacterInput!) {\n    updateCharacter(id: $updateCharacterId, input: $input) {\n      id\n      name\n      age\n      level\n      race {\n        id\n        indx\n        name\n      }\n      class {\n        id\n        indx\n        name\n      }\n    }\n  }\n"): (typeof documents)["\n  mutation UpdateCharacter($updateCharacterId: ID!, $input: UpdateCharacterInput!) {\n    updateCharacter(id: $updateCharacterId, input: $input) {\n      id\n      name\n      age\n      level\n      race {\n        id\n        indx\n        name\n      }\n      class {\n        id\n        indx\n        name\n      }\n    }\n  }\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;