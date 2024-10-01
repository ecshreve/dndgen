import { useQuery } from "@apollo/client";
import InfoIcon from "@mui/icons-material/Info";
import { Box, IconButton, Input, Tooltip, Typography } from "@mui/material";
import React, { useState } from "react";
import { GET_ABILITY_SCORE_DETAILS } from "../../queries/getAbilityScoreDetails";
import "./builder.css";
interface AbilityScore {
  indx: string;
  name: string;
  fullName: string;
  desc: string;
}

interface AbilityScoreData {
  abilityScores: AbilityScore[];
}

const AbilityScorePicker= () => {
  const { data, loading, error } = useQuery<AbilityScoreData>(
    GET_ABILITY_SCORE_DETAILS
  );
  const [showDesc, setShowDesc] = useState<{ [key: string]: boolean }>({});

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  const abilityScores = data?.abilityScores || [];
  const [scoreValues, setScoreValues] = useState<{ [key: string]: number }>({
    str: 8,
    dex: 8,
    con: 8,
    int: 8,
    wis: 8,
    cha: 8,
  });

  const handleScoreChange = (indx: string, value: string) => {
    setScoreValues((prev) => ({ ...prev, [indx]: parseInt(value, 10) }));
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
  }

  return (
    <Box marginTop={2} marginBottom={2}>
      <Typography variant="h4">Ability Score Picker</Typography>
      <Box display="flex" flexDirection="row" alignItems="center" flexWrap="wrap" justifyContent="space-between">
        {abilityScores.map((score) => (
          <Box key={score.indx} display="flex" flexDirection="row" alignItems="center" gap={1} border="1px solid black" borderRadius={2} padding={2}>
            <Box display="flex" flexDirection="column" alignItems="center" gap={1}>
              <Box display="flex" flexDirection="row" alignItems="center" gap={1}>
                <Box>
                  <Typography>{score.name}</Typography>
                </Box>
                <Box>
                  <Tooltip title={score.desc}>
                    <IconButton size="small">
                      <InfoIcon />
                    </IconButton>
                  </Tooltip>
                </Box>
              </Box>
              <Box display="flex" flexDirection="row" alignItems="center" gap={1}>
                <Box width="50px">
                  <Input
                    type="number"
                    value={scoreValues[score.indx]}
                    onChange={(e) => handleScoreChange(score.indx, e.target.value)}
                      />
                </Box>
                <Box className="ability-modifier" sx={{backgroundColor: abilityScoreModifier(scoreValues[score.indx]) >= 0 ? "lightgreen" : "lightcoral"}}>{abilityScoreModifier(scoreValues[score.indx])}</Box>
              </Box>
            </Box>
            
          </Box>
        ))}
      </Box>
    </Box>
  );
};

export default AbilityScorePicker;
