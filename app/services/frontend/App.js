import React from 'react';
import { StatusBar } from 'expo-status-bar';
import { StyleSheet, TextInput, View, Text, ScrollView } from 'react-native';
import { CALC_API } from '@env';

export default function App() {
    const [number, onChangeNumber] = React.useState('');
    const [error, setError] = React.useState('');
    const [response, setResponse] = React.useState({
        total: null,
        sumGroup: [],
    });

    function execute(event) {
        fetch(`${CALC_API}execute`, {
            method: 'POST',
            headers: {
                Accept: 'application/json',
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                input: event.nativeEvent.text,
            }),
        })
            .then((r) => {
                return r.json();
            })
            .then((r) => {
                setResponse(r);
            })
            .catch((e) => {
                setError(e.message);
            });
    }

    return (
        <View style={styles.container}>
            <Text style={styles.calc}>CALC</Text>
            <TextInput
                style={styles.input}
                onChangeText={onChangeNumber}
                value={number}
                keyboardType="phone-pad"
                placeholder="Enter text:"
                onSubmitEditing={execute}
            />
            {error.length ? <Text> {error} </Text> : null}
            <ScrollView style={styles.resultContainer}>
                {response?.sumGroup?.length
                    ? response.sumGroup.map((e, i) => {
                          return (
                              <React.Fragment key={i + 1}>
                                  <Text style={styles.resultText}>
                                      Group {i + 1}: {e}
                                  </Text>
                              </React.Fragment>
                          );
                      })
                    : undefined}
                {response?.total != null ? (
                    <Text style={styles.resultText}>
                        Total: {response.total}
                    </Text>
                ) : undefined}
            </ScrollView>
            <StatusBar style="auto" />
        </View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#eee',
        alignItems: 'center',
        justifyContent: 'start',
        paddingTop: 20,
    },
    resultContainer: {
        width: 800,
        maxWidth: '90%',
        margin: 10,
    },
    resultText: {
        fontSize: 20,
        marginTop: 4,
        fontWeight: '500',
    },
    input: {
        height: 40,
        width: 800,
        maxWidth: '90%',
        margin: 12,
        borderWidth: 1,
        padding: 10,
    },
    calc: {
        marginBottom: 50,
        marginTop: 50,
        fontWeight: '600',
        textAlign: 'center',
        fontSize: 80,
    },
});
