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
    "\n  fragment ProficiencyDetails on Proficiency {\n    id\n    indx\n    reference\n  }\n": types.ProficiencyDetailsFragmentDoc,
    "\n  fragment ProficiencyOptionDetails on ProficiencyChoice {\n    id\n    desc\n    choose\n    proficiencies {\n      ...ProficiencyDetails\n    }\n  }\n": types.ProficiencyOptionDetailsFragmentDoc,
    "\n  fragment EquipmentDetails on Equipment {\n    id\n    indx\n    name\n    weight\n  }\n": types.EquipmentDetailsFragmentDoc,
    "\n  fragment StartingEquipmentDetails on EquipmentEntry {\n    id\n    quantity\n    equipment {\n      ...EquipmentDetails\n    }\n  }\n": types.StartingEquipmentDetailsFragmentDoc,
    "\n  query GetAbilityScoreDetails {\n    abilityScores {\n      indx\n      name\n      fullName\n      desc\n    }\n  }\n": types.GetAbilityScoreDetailsDocument,
    "\n  query GetAbilityScores($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          characterAbilityScores {\n            id\n            score\n            modifier\n            abilityScore {\n              id\n              indx\n              name\n            }\n            characterSkills {\n              id\n              proficient\n              skill {\n                id\n                indx\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n": types.GetAbilityScoresDocument,
    "\n  query GetCharacter($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          name\n          age\n          level\n          proficiencyBonus\n          class {\n            id\n            indx\n            name\n          }\n          race {\n            id\n            indx\n            name\n          }\n          alignment {\n            id\n            indx\n            name\n          }\n          characterAbilityScores {\n            id\n            score\n            modifier\n            abilityScore {\n              id\n              indx\n              name\n            }\n            characterSkills {\n              id\n              proficient\n              skill {\n                id\n                indx\n                name\n              }\n            }\n          }\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n              }\n            }\n          }\n        }\n      }\n    }\n  }\n": types.GetCharacterDocument,
    "\n  \n  \n  \n  fragment ClassDetails on Class {\n    id\n    indx\n    name\n    savingThrows {\n      id\n      indx\n    }\n    proficiencies {\n      ...ProficiencyDetails\n    }\n    proficiencyOptions {\n      ...ProficiencyOptionDetails\n    }\n    startingEquipment {\n      ...StartingEquipmentDetails\n    }\n  }\n\n  query GetClassDetails {\n    classes {\n      ...ClassDetails\n    }\n  }\n": types.ClassDetailsFragmentDoc,
    "\n  query GetClasses {\n    classes {\n      id\n      indx\n      name\n    }\n  }\n": types.GetClassesDocument,
    "\n  \n  \n  fragment RaceDetails on Race {\n    id\n    indx\n    name\n    ageDesc\n    alignmentDesc\n    size\n    sizeDesc\n    speed\n    languageDesc\n    startingProficiencies {\n      ...ProficiencyDetails\n    }\n    startingProficiencyOptions {\n      ...ProficiencyOptionDetails\n    }\n  }\n\n  query GetRaceDetails {\n    races {\n      ...RaceDetails\n    }\n  }\n": types.RaceDetailsFragmentDoc,
    "\n  query GetRaces {\n    races {\n      id\n      indx\n      name\n    }\n  }\n": types.GetRacesDocument,
    "\n  query GetSkillDetails {\n    skills {\n      id\n      indx\n      name\n      desc\n      abilityScore {\n        id\n        indx\n        name\n      }\n    }\n  }\n": types.GetSkillDetailsDocument,
    "\n  query GetSkills($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          proficiencyBonus\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n            }\n          }\n        }\n      }\n    }\n  }\n}\n": types.GetSkillsDocument,
    "\n  mutation UpdateCharacter($updateCharacterId: ID!, $input: UpdateCharacterInput!) {\n    updateCharacter(id: $updateCharacterId, input: $input) {\n      id\n    }\n  }\n": types.UpdateCharacterDocument,
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
export function gql(source: "\n  fragment ProficiencyDetails on Proficiency {\n    id\n    indx\n    reference\n  }\n"): (typeof documents)["\n  fragment ProficiencyDetails on Proficiency {\n    id\n    indx\n    reference\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  fragment ProficiencyOptionDetails on ProficiencyChoice {\n    id\n    desc\n    choose\n    proficiencies {\n      ...ProficiencyDetails\n    }\n  }\n"): (typeof documents)["\n  fragment ProficiencyOptionDetails on ProficiencyChoice {\n    id\n    desc\n    choose\n    proficiencies {\n      ...ProficiencyDetails\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  fragment EquipmentDetails on Equipment {\n    id\n    indx\n    name\n    weight\n  }\n"): (typeof documents)["\n  fragment EquipmentDetails on Equipment {\n    id\n    indx\n    name\n    weight\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  fragment StartingEquipmentDetails on EquipmentEntry {\n    id\n    quantity\n    equipment {\n      ...EquipmentDetails\n    }\n  }\n"): (typeof documents)["\n  fragment StartingEquipmentDetails on EquipmentEntry {\n    id\n    quantity\n    equipment {\n      ...EquipmentDetails\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetAbilityScoreDetails {\n    abilityScores {\n      indx\n      name\n      fullName\n      desc\n    }\n  }\n"): (typeof documents)["\n  query GetAbilityScoreDetails {\n    abilityScores {\n      indx\n      name\n      fullName\n      desc\n    }\n  }\n"];
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
export function gql(source: "\n  \n  \n  \n  fragment ClassDetails on Class {\n    id\n    indx\n    name\n    savingThrows {\n      id\n      indx\n    }\n    proficiencies {\n      ...ProficiencyDetails\n    }\n    proficiencyOptions {\n      ...ProficiencyOptionDetails\n    }\n    startingEquipment {\n      ...StartingEquipmentDetails\n    }\n  }\n\n  query GetClassDetails {\n    classes {\n      ...ClassDetails\n    }\n  }\n"): (typeof documents)["\n  \n  \n  \n  fragment ClassDetails on Class {\n    id\n    indx\n    name\n    savingThrows {\n      id\n      indx\n    }\n    proficiencies {\n      ...ProficiencyDetails\n    }\n    proficiencyOptions {\n      ...ProficiencyOptionDetails\n    }\n    startingEquipment {\n      ...StartingEquipmentDetails\n    }\n  }\n\n  query GetClassDetails {\n    classes {\n      ...ClassDetails\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetClasses {\n    classes {\n      id\n      indx\n      name\n    }\n  }\n"): (typeof documents)["\n  query GetClasses {\n    classes {\n      id\n      indx\n      name\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  \n  \n  fragment RaceDetails on Race {\n    id\n    indx\n    name\n    ageDesc\n    alignmentDesc\n    size\n    sizeDesc\n    speed\n    languageDesc\n    startingProficiencies {\n      ...ProficiencyDetails\n    }\n    startingProficiencyOptions {\n      ...ProficiencyOptionDetails\n    }\n  }\n\n  query GetRaceDetails {\n    races {\n      ...RaceDetails\n    }\n  }\n"): (typeof documents)["\n  \n  \n  fragment RaceDetails on Race {\n    id\n    indx\n    name\n    ageDesc\n    alignmentDesc\n    size\n    sizeDesc\n    speed\n    languageDesc\n    startingProficiencies {\n      ...ProficiencyDetails\n    }\n    startingProficiencyOptions {\n      ...ProficiencyOptionDetails\n    }\n  }\n\n  query GetRaceDetails {\n    races {\n      ...RaceDetails\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetRaces {\n    races {\n      id\n      indx\n      name\n    }\n  }\n"): (typeof documents)["\n  query GetRaces {\n    races {\n      id\n      indx\n      name\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetSkillDetails {\n    skills {\n      id\n      indx\n      name\n      desc\n      abilityScore {\n        id\n        indx\n        name\n      }\n    }\n  }\n"): (typeof documents)["\n  query GetSkillDetails {\n    skills {\n      id\n      indx\n      name\n      desc\n      abilityScore {\n        id\n        indx\n        name\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetSkills($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          proficiencyBonus\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"): (typeof documents)["\n  query GetSkills($characterId: ID!) {\n    characters(where: {id: $characterId}) {\n      edges {\n        node {\n          id\n          proficiencyBonus\n          characterSkills {\n            id\n            proficient\n            skill {\n              id\n              indx\n              name\n            }\n            characterAbilityScore {\n              id\n              score\n              modifier\n              abilityScore {\n                id\n                indx\n                name\n            }\n          }\n        }\n      }\n    }\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  mutation UpdateCharacter($updateCharacterId: ID!, $input: UpdateCharacterInput!) {\n    updateCharacter(id: $updateCharacterId, input: $input) {\n      id\n    }\n  }\n"): (typeof documents)["\n  mutation UpdateCharacter($updateCharacterId: ID!, $input: UpdateCharacterInput!) {\n    updateCharacter(id: $updateCharacterId, input: $input) {\n      id\n    }\n  }\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;