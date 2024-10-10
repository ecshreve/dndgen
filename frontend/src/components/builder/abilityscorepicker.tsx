import { useQuery } from "@apollo/client";
import InfoIcon from "@mui/icons-material/Info";
import { Box, Grid2, IconButton, Tooltip, Typography } from "@mui/material";
import React from "react";
import { GET_ABILITY_SCORE_DETAILS } from "../../queries/getAbilityScoreDetails";
import { abilityScoreToModifier } from "../../utils";
import "./builder.css";
import NumberInput from "./customnumberinput";
interface AbilityScore {
  indx: string;
  name: string;
  fullName: string;
  desc: string;
}

interface AbilityScoreData {
  abilityScores: AbilityScore[];
}

interface AbilityScorePickerProps {
  scoreValues: { [key: string]: number };
  handleChange: (indx: string, value: number) => void;
  enableEdit: boolean;
}
const AbilityScorePicker = ({ scoreValues, handleChange, enableEdit }: AbilityScorePickerProps) => {
  

  const { data, loading, error } = useQuery<AbilityScoreData>(
    GET_ABILITY_SCORE_DETAILS
  );

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  const abilityScores = data?.abilityScores || [];

  return (
    <Box sx={{ maxWidth: '1000px', margin: '0 auto', padding: '20px', border: '1px solid #ccc', borderRadius: '10px', backgroundColor: '#f9f9f9' }}>
      <Typography variant="h5" gutterBottom borderBottom="1px solid #ccc">
        Ability Scores
      </Typography>
      <Grid2 container spacing={1}>
        {abilityScores.map((score) => (
          <Grid2
            key={score.indx}
            size={{xs: 6, sm: 4, md: 2, lg: 2}}
            display="flex"
            flexDirection="column"
            alignItems="center"
          >
            <Box
              display="flex"
              flexDirection="column"
              alignItems="center"
              border="1px solid black"
              borderRadius={2}
              padding={1}
              gap={2}
            >
              <Box display="flex" flexDirection="row" alignItems="center" gap={1}>
                <Tooltip title={score.desc}>
                  <IconButton size="small">
                    <InfoIcon />
                    <Typography variant="h6">{score.name}</Typography>
                  </IconButton>
                </Tooltip>
                <div
                  className="ability-modifier"
                  style={{
                    backgroundColor:
                      abilityScoreToModifier(scoreValues[score.indx]) >= 0
                        ? "lightgreen"
                        : "lightcoral",
                  }}
                >
                  {abilityScoreToModifier(scoreValues[score.indx]) >= 0 ? "+" : ""}
                  {abilityScoreToModifier(scoreValues[score.indx])}
                </div>
              </Box>
              {enableEdit ? (
                <NumberInput
                  min={1}
                  max={30}
                  value={scoreValues[score.indx]}
                  onChange={(event, newValue) =>
                    handleChange(score.indx, newValue ?? 0)
                  }
                />
              ) : (
                <Typography variant="h6">{scoreValues[score.indx]}</Typography>
              )}
            </Box>
          </Grid2>
        ))}
      </Grid2>
    </Box>
  );
};

export default AbilityScorePicker;
