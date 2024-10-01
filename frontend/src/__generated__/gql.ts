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
    "\n  mutation CreateCharacter($input: CreateCharacterInput!) {\n    createCharacter(input: $input) {\n      id\n      name\n      age\n      level\n      race {\n        id\n        indx\n        name\n      }\n      class {\n        id\n        indx\n        name\n      }\n    }\n  }\n": types.CreateCharacterDocument,
    "\n  query GetAbilityScores($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          characterAbilityScores {\n            id\n            score\n            modifier\n            abilityScore {\n              id\n              indx\n              name\n            }\n            characterSkills {\n              id\n              proficient\n              skill {\n                id\n                indx\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n": types.GetAbilityScoresDocument,
    "\n  query GetCharacter($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          name\n          age\n          level\n          proficiencyBonus\n          class {\n            id\n            indx\n            name\n          }\n          race {\n            id\n            indx\n            name\n          }\n          alignment {\n            id\n            indx\n            name\n          }\n          characterAbilityScores {\n            id\n            score\n            modifier\n            abilityScore {\n              id\n              indx\n              name\n            }\n            characterSkills {\n              id\n              proficient\n              skill {\n                id\n                indx\n                name\n              }\n            }\n          }\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n": types.GetCharacterDocument,
    "\n  query GetClassDetails {\n    classes {\n      id\n      indx\n      name\n      savingThrows {\n        id\n        indx\n        name\n      }\n      proficiencies {\n        id\n        indx\n        name\n      }\n      proficiencyOptions {\n        id\n        desc\n      }\n      startingEquipment {\n        id\n        quantity\n      }\n    }\n  }\n": types.GetClassDetailsDocument,
    "\n  query GetClasses {\n    classes {\n      id\n      indx\n      name\n    }\n  }\n": types.GetClassesDocument,
    "\n  query GetRaceDetails {\n    races {\n      id\n      indx\n      name\n      ageDesc\n      alignmentDesc\n      size\n      sizeDesc\n      speed\n      languageDesc\n      startingProficiencies {\n        id\n        indx\n        name\n      }\n      startingProficiencyOptions {\n        id\n        desc\n        choose\n        proficiencies {\n          id\n          indx\n          name\n        }\n      }\n    }\n  }\n": types.GetRaceDetailsDocument,
    "\n  query GetRaces {\n    races {\n      id\n      indx\n      name\n    }\n  }\n": types.GetRacesDocument,
    "\n  query GetSkills($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          proficiencyBonus\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n            }\n          }\n        }\n      }\n    }\n  }\n}\n": types.GetSkillsDocument,
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
export function gql(source: "\n  mutation CreateCharacter($input: CreateCharacterInput!) {\n    createCharacter(input: $input) {\n      id\n      name\n      age\n      level\n      race {\n        id\n        indx\n        name\n      }\n      class {\n        id\n        indx\n        name\n      }\n    }\n  }\n"): (typeof documents)["\n  mutation CreateCharacter($input: CreateCharacterInput!) {\n    createCharacter(input: $input) {\n      id\n      name\n      age\n      level\n      race {\n        id\n        indx\n        name\n      }\n      class {\n        id\n        indx\n        name\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetAbilityScores($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          characterAbilityScores {\n            id\n            score\n            modifier\n            abilityScore {\n              id\n              indx\n              name\n            }\n            characterSkills {\n              id\n              proficient\n              skill {\n                id\n                indx\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query GetAbilityScores($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          characterAbilityScores {\n            id\n            score\n            modifier\n            abilityScore {\n              id\n              indx\n              name\n            }\n            characterSkills {\n              id\n              proficient\n              skill {\n                id\n                indx\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetCharacter($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          name\n          age\n          level\n          proficiencyBonus\n          class {\n            id\n            indx\n            name\n          }\n          race {\n            id\n            indx\n            name\n          }\n          alignment {\n            id\n            indx\n            name\n          }\n          characterAbilityScores {\n            id\n            score\n            modifier\n            abilityScore {\n              id\n              indx\n              name\n            }\n            characterSkills {\n              id\n              proficient\n              skill {\n                id\n                indx\n                name\n              }\n            }\n          }\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query GetCharacter($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          name\n          age\n          level\n          proficiencyBonus\n          class {\n            id\n            indx\n            name\n          }\n          race {\n            id\n            indx\n            name\n          }\n          alignment {\n            id\n            indx\n            name\n          }\n          characterAbilityScores {\n            id\n            score\n            modifier\n            abilityScore {\n              id\n              indx\n              name\n            }\n            characterSkills {\n              id\n              proficient\n              skill {\n                id\n                indx\n                name\n              }\n            }\n          }\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetClassDetails {\n    classes {\n      id\n      indx\n      name\n      savingThrows {\n        id\n        indx\n        name\n      }\n      proficiencies {\n        id\n        indx\n        name\n      }\n      proficiencyOptions {\n        id\n        desc\n      }\n      startingEquipment {\n        id\n        quantity\n      }\n    }\n  }\n"): (typeof documents)["\n  query GetClassDetails {\n    classes {\n      id\n      indx\n      name\n      savingThrows {\n        id\n        indx\n        name\n      }\n      proficiencies {\n        id\n        indx\n        name\n      }\n      proficiencyOptions {\n        id\n        desc\n      }\n      startingEquipment {\n        id\n        quantity\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetClasses {\n    classes {\n      id\n      indx\n      name\n    }\n  }\n"): (typeof documents)["\n  query GetClasses {\n    classes {\n      id\n      indx\n      name\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetRaceDetails {\n    races {\n      id\n      indx\n      name\n      ageDesc\n      alignmentDesc\n      size\n      sizeDesc\n      speed\n      languageDesc\n      startingProficiencies {\n        id\n        indx\n        name\n      }\n      startingProficiencyOptions {\n        id\n        desc\n        choose\n        proficiencies {\n          id\n          indx\n          name\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query GetRaceDetails {\n    races {\n      id\n      indx\n      name\n      ageDesc\n      alignmentDesc\n      size\n      sizeDesc\n      speed\n      languageDesc\n      startingProficiencies {\n        id\n        indx\n        name\n      }\n      startingProficiencyOptions {\n        id\n        desc\n        choose\n        proficiencies {\n          id\n          indx\n          name\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetRaces {\n    races {\n      id\n      indx\n      name\n    }\n  }\n"): (typeof documents)["\n  query GetRaces {\n    races {\n      id\n      indx\n      name\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetSkills($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          proficiencyBonus\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"): (typeof documents)["\n  query GetSkills($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          proficiencyBonus\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation UpdateCharacter($updateCharacterId: ID!, $input: UpdateCharacterInput!) {\n    updateCharacter(id: $updateCharacterId, input: $input) {\n      id\n      name\n      age\n      level\n      race {\n        id\n        indx\n        name\n      }\n      class {\n        id\n        indx\n        name\n      }\n    }\n  }\n"): (typeof documents)["\n  mutation UpdateCharacter($updateCharacterId: ID!, $input: UpdateCharacterInput!) {\n    updateCharacter(id: $updateCharacterId, input: $input) {\n      id\n      name\n      age\n      level\n      race {\n        id\n        indx\n        name\n      }\n      class {\n        id\n        indx\n        name\n      }\n    }\n  }\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;