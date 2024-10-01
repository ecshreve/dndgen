import { useQuery } from "@apollo/client";
import InfoIcon from "@mui/icons-material/Info";
import { Box, Grid2, IconButton, Tooltip, Typography } from "@mui/material";
import React, { useState } from "react";
import { GET_ABILITY_SCORE_DETAILS } from "../../queries/getAbilityScoreDetails";
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

const AbilityScorePicker = ({ enableEdit }: { enableEdit: boolean }) => {
  const [scoreValues, setScoreValues] = useState<{ [key: string]: number }>({
    str: 8,
    dex: 8,
    con: 8,
    int: 8,
    wis: 8,
    cha: 8,
  });

  const { data, loading, error } = useQuery<AbilityScoreData>(
    GET_ABILITY_SCORE_DETAILS
  );

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  const abilityScores = data?.abilityScores || [];

  const handleScoreChange = (indx: string, value: number) => {
    setScoreValues((prev) => ({ ...prev, [indx]: value }));
  };

  const abilityScoreModifier = (score: number): number => {
    if (score < 1) {
      return -5;
    }

    if (score > 30) {
      return 10;
    }

    const mod = Math.floor((score - 10) / 2);

    return mod;
  };

  return (
    <Grid2 container spacing={2} margin={2}>
      <Grid2 size={12} display="flex" justifyContent="center">
        <Typography variant="h4">Ability Scores</Typography>
      </Grid2>
      {abilityScores.map((score) => (
        <Grid2
          size={{ xs: 12, sm: 6, md: 4, lg: 3 }}
          display="flex"
          flexDirection="column"
          alignItems="center"
        >
          <Box
            key={score.indx}
            display="flex"
            flexDirection="column"
            alignItems="center"
            border="1px solid black"
            borderRadius={2}
            padding={1}
            gap={2}
          >
            <Box display="flex" flexDirection="row" alignItems="center" gap={4}>
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
                    abilityScoreModifier(scoreValues[score.indx]) >= 0
                      ? "lightgreen"
                      : "lightcoral",
                }}
              >
                {abilityScoreModifier(scoreValues[score.indx]) >= 0 ? "+" : ""}
                {abilityScoreModifier(scoreValues[score.indx])}
              </div>
            </Box>
            {enableEdit ? (
              <NumberInput
                min={1}
                max={30}
                value={scoreValues[score.indx]}
                onChange={(event, newValue) =>
                  handleScoreChange(score.indx, newValue ?? 0)
                }
              />
            ) : (
              <Typography variant="h6">{scoreValues[score.indx]}</Typography>
            )}
          </Box>
        </Grid2>
      ))}
    </Grid2>
  );
};

export default AbilityScorePicker;
