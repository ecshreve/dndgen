import { gql } from "@apollo/client";

export const CREATE_CHARACTER = gql`
  mutation CreateCharacter($input: CreateCharacterInput!) {
    createCharacter(input: $input) {
      id
      name
      age
      level
      race {
        id
        indx
        name
      }
      class {
        id
        indx
        name
      }
    }
  }
`;
