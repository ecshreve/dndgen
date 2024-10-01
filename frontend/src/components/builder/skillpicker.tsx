import { useQuery } from "@apollo/client";
import CheckCircleIcon from "@mui/icons-material/CheckCircle";
import RadioButtonUncheckedIcon from "@mui/icons-material/RadioButtonUnchecked";
import {
  Box,
  IconButton,
  Paper,
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Typography,
} from "@mui/material";
import React from "react";
import { GET_SKILL_DETAILS } from "../../queries/getSkillDetails";
import { abilityScoreToModifier } from "../../utils";

interface SkillPickerProps {
  abilityScoreValues: { [key: string]: number };
  proficiencies: string[];
  profBonus: number;
  enableEdit: boolean;
}
const SkillPicker = ({
  abilityScoreValues,
  proficiencies,
  profBonus,
  enableEdit,
}: SkillPickerProps) => {
  const { data, loading, error } = useQuery(GET_SKILL_DETAILS);
  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  const calculateSkillModifier = (skillId: string) => {
    const skill = data.skills.find((skill: any) => skill.id === skillId);
    const asMod = abilityScoreToModifier(
      abilityScoreValues[skill.abilityScore.indx]
    );
    const profMod =
      proficiencies.includes(`skill-${skill.indx}`) ||
      proficiencies.includes(`${skill.indx}`)
        ? profBonus
        : 0;
    return asMod + profMod;
  };

  return (
    <Box
      sx={{
        maxWidth: "1000px",
        margin: "0 auto",
        padding: "20px",
        border: "1px solid #ccc",
        borderRadius: "10px",
        backgroundColor: "#f9f9f9",
      }}
    >
      <Typography
        variant="h5"
        gutterBottom
        sx={{ borderBottom: "1px solid #ccc" }}
      >
        Skills
      </Typography>
      <TableContainer component={Paper}>
        <Table size="small">
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
                  {abilityScoreToModifier(
                    abilityScoreValues[skill.abilityScore.indx]
                  ) >= 0
                    ? "+"
                    : ""}
                  {abilityScoreToModifier(
                    abilityScoreValues[skill.abilityScore.indx]
                  )}
                </TableCell>
                <TableCell>
                  <IconButton disabled={!enableEdit}>
                    {proficiencies.includes(`skill-${skill.indx}`) ||
                    proficiencies.includes(`${skill.indx}`) ? (
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
