import { useQuery } from '@apollo/client';
import CheckCircleIcon from '@mui/icons-material/CheckCircle';
import RadioButtonUncheckedIcon from '@mui/icons-material/RadioButtonUnchecked';
import { Box, IconButton, Paper, Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Typography } from '@mui/material';
import React, { useState } from 'react';
import { GET_SKILL_DETAILS } from '../../queries/getSkillDetails';
import { abilityScoreToModifier } from '../../utils';

interface SkillPickerProps {
    abilityScoreValues: { [key: string]: number };
    proficientSkills: string[];
    enableEdit: boolean;
}
const SkillPicker = ({ abilityScoreValues, proficientSkills, enableEdit }: SkillPickerProps) => {
  const [selectedSkills, setSelectedSkills] = useState<string[]>([]);

  const {data, loading, error} = useQuery(GET_SKILL_DETAILS);
  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  const isSkillSelected = (skillId: string) => {
    return selectedSkills.includes(skillId);
  };

  const calculateSkillModifier = (skillId: string) => {
    const skill = data.skills.find((skill: any) => skill.id === skillId);
    const asMod = abilityScoreToModifier(abilityScoreValues[skill.abilityScore.indx]);
    const profMod = proficientSkills.includes(`skill-${skill.indx}`) ? 2 : 0;
    return asMod + profMod;
  };

  console.log(proficientSkills);
  return (
    <Box sx={{ maxWidth: '1000px', margin: '0 auto', padding: '20px', border: '1px solid #ccc', borderRadius: '10px', backgroundColor: '#f9f9f9' }}>
      <Typography variant="h5" gutterBottom sx={{ borderBottom: '1px solid #ccc'}}>
        Skills
      </Typography>
      <TableContainer component={Paper}>
        <Table size='small'>
          <TableHead>
            <TableRow>
              <TableCell>Skill</TableCell>
              <TableCell>Ability Score</TableCell>
              <TableCell>Ability Score Modifier</TableCell>
              <TableCell>Proficient</TableCell>
              <TableCell>Skill Modifier</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {data.skills.map((skill: any) => (
              <TableRow key={skill.id}>
                <TableCell>{skill.name}</TableCell>
                <TableCell>
                    <Box mr={1}>{skill.abilityScore.indx}</Box>
                </TableCell>
                <TableCell>
                  {abilityScoreToModifier(abilityScoreValues[skill.abilityScore.indx]) >= 0 ? "+" : ""}
                  {abilityScoreToModifier(abilityScoreValues[skill.abilityScore.indx])}
                </TableCell>
                <TableCell>
                  <IconButton disabled={!enableEdit}>
                    {isSkillSelected(skill.id) ? (
                      <CheckCircleIcon color="primary" />
                    ) : proficientSkills.includes(`skill-${skill.indx}`) ? (
                      <CheckCircleIcon color="secondary" />
                    ) : (
                      <RadioButtonUncheckedIcon color="primary" />
                    )}
                  </IconButton>
                </TableCell>
                <TableCell>
                  {calculateSkillModifier(skill.id) >= 0 ? "+" : ""}
                  {calculateSkillModifier(skill.id)}
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
      
    </Box>
  );
};

export default SkillPicker;