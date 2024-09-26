import Table from '@mui/material/Table';
import TableBody from '@mui/material/TableBody';
import TableCell from '@mui/material/TableCell';
import TableHead from '@mui/material/TableHead';
import TableRow from '@mui/material/TableRow';

import { useQuery } from '@apollo/client';

import { gql } from './__generated__/gql';

const GET_CHARACTERS = gql(/* GraphQL */ `
  query GetCharacters {
    characters {
      edges {
        node {
          id
          name
          race {
            id
            indx
          }
          class {
            id
            indx
          }
        }
      }
    }
  }
`);

const CharacterList = () => {
  // our query's result, data, is typed!
  const { loading, data } = useQuery(
    GET_CHARACTERS,
    {}
  );
  return (
    <div>
      <h3>All Characters</h3>
      {loading ? (
        <p>Loading ...</p>
      ) : (
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>id</TableCell>
              <TableCell>name</TableCell>
              <TableCell>race</TableCell>
              <TableCell>class</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data && data.characters.edges?.map(character => (
              <TableRow>
                <TableCell>{character?.node?.id}</TableCell>
                <TableCell>{character?.node?.name}</TableCell>
                <TableCell>{character?.node?.race.indx}</TableCell>
                <TableCell>{character?.node?.class.indx}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      )}
    </div>
  );
}

export default CharacterList;