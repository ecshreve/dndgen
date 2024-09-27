import { gql } from '@apollo/client';

export const GET_CLASSES = gql`
  query GetClasses {
    classes {
      id
      indx
      name
    }
  }
`;
