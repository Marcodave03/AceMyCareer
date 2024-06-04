import 'dart:convert';
import 'dart:io';
import 'dart:typed_data';
import 'package:audioplayers/audioplayers.dart';
import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class TtsTesting extends StatefulWidget {
  const TtsTesting({super.key});

  @override
  _TtsTestingState createState() => _TtsTestingState();
}

class _TtsTestingState extends State<TtsTesting> {
  final TextEditingController _textController = TextEditingController();
  AudioPlayer _audioPlayer = AudioPlayer();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        body: Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        TextField(
          controller: _textController,
          decoration: const InputDecoration(
            labelText: 'Enter Text',
          ),
        ),
        const SizedBox(height: 20),
        ElevatedButton(
          onPressed: () async {
            String inputText = _textController.text;
            if (inputText.isNotEmpty) {
              await getTextToSpeech(inputText);
              print('Text-to-speech conversion completed!');
            } else {
              print('Please enter text before pressing the button.');
            }
          },
          child: const Text('Call Flask Endpoint'),
        ),
      ],
    ));
  }

  Future<void> getTextToSpeech(inputText) async {
    final String serverUrl = 'http://10.0.2.2:5000/send-text/';

    final response = await http.get(Uri.parse('$serverUrl$inputText'));
    if (response.statusCode == 202) {
      final Map<String, dynamic> jsonResponse = json.decode(response.body);
      final String jobId = jsonResponse['job_id'];
      print('Job accepted with ID: $jobId');
      String audioUrl = 'http://10.0.2.2:5000/get-audio/$jobId';
      _audioPlayer.play(UrlSource(audioUrl));
    } else {
      // Handle error
      print('Error: ${response.body}');
    }
  }

  @override
  void dispose() {
    _textController.dispose();
    super.dispose();
  }
}
