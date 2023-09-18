import React from "react";
import { StatusBar } from "expo-status-bar";
import { StyleSheet, TextInput, View, Text } from "react-native";

export default function App() {
  const [number, onChangeNumber] = React.useState("");
  const [result, setResult] = React.useState("");

  function execute(event) {
    setResult(event.nativeEvent.text);
  }

  return (
    <View style={styles.container}>
      <TextInput
        style={styles.input}
        onChangeText={onChangeNumber}
        value={number}
        placeholder="Enter text:"
        onSubmitEditing={execute}
      />
      <Text>{result}</Text>
      <StatusBar style="auto" />
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "start",
    marginTop: 100,
  },
  input: {
    height: 40,
    width: 100,
    margin: 12,
    borderWidth: 1,
    padding: 10,
  },
});
