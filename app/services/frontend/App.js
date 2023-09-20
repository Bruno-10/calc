import React from "react";
import { StatusBar } from "expo-status-bar";
import { StyleSheet, TextInput, View, Text } from "react-native";
const API_URL = "http://192.168.1.91:3000/v1";

export default function App() {
  const [number, onChangeNumber] = React.useState("");
  const [error, setError] = React.useState("");
  const [response, setResponse] = React.useState({ total: null, sumGroup: [] });

  function execute(event) {
    fetch(`${API_URL}/execute`, {
      method: "POST",
      headers: {
        Accept: "application/json",
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        input: event.nativeEvent.text,
      }),
    })
      .then(async (r) => {
        const data = await r.json();
        setResponse(data);
      })
      .catch((error) => {
        console.log(error.stack);
        setError(error.message);
      });
  }

  return (
    <View style={styles.container}>
      <Text style={styles.calc}>
        <Text> CALC </Text>
      </Text>
      <TextInput
        style={styles.input}
        onChangeText={onChangeNumber}
        value={number}
        keyboardType="phone-pad"
        placeholder="Enter text:"
        onSubmitEditing={execute}
      />
      {error.length ? <Text> {error} </Text> : null}
      <View style={styles.resultContainer}>
        {response.total != null ? (
          <Text style={styles.resultText}>Total: {response.total}</Text>
        ) : undefined}
        {response.sumGroup.length
          ? response.sumGroup.map((e, i) => {
              return (
                <Text key={i + 1} style={styles.resultText}>
                  Group {i + 1}: {e}
                </Text>
              );
            })
          : undefined}
      </View>
      <StatusBar style="auto" />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#eee",
    alignItems: "center",
    justifyContent: "start",
    paddingTop: 200,
  },
  resultContainer: {
    width: 800,
    maxWidth: "90%",
    margin: 10,
  },
  resultText: {
    fontSize: "2rem",
    marginTop: 4,
    fontWeight: 500,
  },
  input: {
    height: 40,
    width: 800,
    maxWidth: "90%",
    margin: 12,
    borderWidth: 1,
    padding: 10,
  },
  calc: {
    marginBottom: "100px",
    fontWeight: "600",
    textAlign: "center",
    fontSize: 60,
  },
});
