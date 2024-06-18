import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get/get_core/src/get_main.dart';
import 'package:speech_to_text/speech_recognition_result.dart';
import 'package:speech_to_text/speech_to_text.dart';
import 'package:text_to_speech/text_to_speech.dart';
import 'package:video_player/video_player.dart';
import 'package:flutter_tts/flutter_tts.dart';
import 'package:http/http.dart' as http;

import 'feedback_page.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

void main() {
  runApp(MaterialApp(
    home: RecordPage(),
  ));
}

class RecordPage extends StatefulWidget {
  late final String feedbackTextFromGPT; 

  @override
  RecordPageState createState() => RecordPageState();
}

class RecordPageState extends State<RecordPage> {
  FlutterTts flutterTts = FlutterTts();
  TextToSpeech tts = TextToSpeech();
  bool _isSpeaking = false;

  SpeechToText _speechToText = SpeechToText();
  bool speechEnabled = false;

  String _lastWords = "";
  VideoPlayerController? _controller; //

  @override
  void initState() {
    super.initState();
    _initSpeechToText();
    _initVideoPlayer(); //
    _setupTts();
  }

  void _initSpeechToText() async {
    await _speechToText.initialize();
    setState(() {});
  }

  void _startListening() async {
    await _speechToText.listen(onResult: _onSpeechResult);

    setState(() {});
  }

  Future<void> _stopListening() async {
    await _speechToText.stop();
    setState(() {});
  }

  void _onSpeechResult(SpeechRecognitionResult result) {
    setState(() {
      _lastWords = result.recognizedWords;
    });
  }

  void _initVideoPlayer() {
    _controller = VideoPlayerController.asset('assets/img/interviewerDummy.mp4')
      ..initialize().then((_) {
        setState(() {});
      });
  }

void _setupTts() {
  flutterTts.setStartHandler(() {
    _isSpeaking = true; 
  });

  flutterTts.setCompletionHandler(() {
    setState(() {
      _isSpeaking = false; 
      _controller?.pause(); 
    });
  });
}

  @override
  void dispose() {
    super.dispose();
    _speechToText.stop();
   flutterTts.stop();
    _controller?.dispose();
  }

  Future<String> getGPTtextFeedback(String _lastWords) async {
    final uri = Uri.parse('https://api.openai.com/v1/chat/completions');

    final apiKey =  dotenv.env['OPENAI_API_KEY'] ?? "";

    final headers = {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer $apiKey'
    };


    final requestBody = json.encode({
      "model": "gpt-3.5-turbo",
      "messages": [
        {
          "role": "system",
          "content":
              "You are Alex, a software engineering interviewer. Your job is to interview candidates for software engineering roles. Ask one question at a time related to software engineering and generate response based on the candidate's responses stored in the $_lastWords variable."},
        {"role": "user", "content": _lastWords}
      ],
      "max_tokens": 200,
    });

    final response = await http.post(uri, headers: headers, body: requestBody);

    if (response.statusCode == 200) {
      final responseBody = json.decode(response.body);
      return responseBody['choices'][0]['message']['content'].trim();
    } else {
      throw Exception(
          'Failed to analyze resume with GPT-3.5 Turbo: ${response.body}');
    }
  }

  Future<String> _convertToText() async {
    var textChatFeedbackGPT = await getGPTtextFeedback(_lastWords);

    return textChatFeedbackGPT;
  }

  Future<void> _textToSpeechForGPT(String text) async {
    if (_speechToText.isListening) return;
    String language = 'en-US';
    tts.setLanguage(language); 
    await tts.speak(text); 
    _controller?.seekTo(Duration.zero);
    _controller?.play(); 
  }

  bool hasSpoken = false;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      appBar: AppBar(
        leading: IconButton(
          icon: const Icon(
            Icons.arrow_back_ios_new,
            color: Color(0xFF2962FF),
          ),
          onPressed: () {
            Get.toNamed("/interviewPage");
          },
        ),
        title: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisSize: MainAxisSize.min,
          children: <Text>[
            Text(
              'Interview Session',
              style: TextStyle(
                  fontWeight: FontWeight.bold,
                  fontSize: 22,
                  color: Color(0xFF2962FF)),
            ),
            Text(
              'Software Engineer', 
              style: TextStyle(
                  fontWeight: FontWeight.bold,
                  fontSize: 16,
                  color: Colors.black), 
            ),
          ],
        ),
      ),
      body: Column(
        
        children: [
          Expanded(
            flex: 2,
            child: Padding(
              padding: const EdgeInsets.all(8.0),
              child: _controller != null && _controller!.value.isInitialized
                  ? AspectRatio(
                      aspectRatio: _controller!.value.aspectRatio,
                      child: VideoPlayer(_controller!),
                    )
                  : Container(
                      alignment: Alignment.center,
                      child: CircularProgressIndicator(),
                    ),
              ),
          ),
          Container(
            color: Colors.white,
            child: ListTile(
              title: Text(
                'Interviewer:',
                style:
                    TextStyle(color: Colors.black, fontWeight: FontWeight.bold),
              ),
            ),
          ),
          Expanded(
            child: Container(
              padding: EdgeInsets.all(10.0),
              margin: EdgeInsets.only(left: 10.0, right: 10.0, bottom: 20.0),
              decoration: BoxDecoration(
                border: Border.all(
                  color: const Color(0xFF2962FF),
                  width: 2.0, 
                ),
                borderRadius: BorderRadius.circular(
                    12), 
              ),
              child: SingleChildScrollView(
                child: FutureBuilder<String>(
                  future: _convertToText(),
                  builder:
                      (BuildContext context, AsyncSnapshot<String> snapshot) {
                    if (snapshot.connectionState == ConnectionState.done) {
                      if (snapshot.hasData) {
                        _textToSpeechForGPT(snapshot.data!);
                        return Text(
                          snapshot.data!, 
                          style: TextStyle(color: Colors.black),
                        );
                      } else if (snapshot.hasError) {
                        return Text(
                          'Error: ${snapshot.error}',
                          style: TextStyle(color: Colors.red),
                        );
                      }
                    }
                    return CircularProgressIndicator();
                  },
                ),
              ),
            ),
          ),
          Container(
            padding: const EdgeInsets.symmetric(horizontal: 16.0),
            margin: EdgeInsets.only(
              top: 10,
              left: 10.0,
              right: 10.0,
            ),
            decoration: BoxDecoration(
                color: Color.fromARGB(255, 14, 35, 154),
                borderRadius: BorderRadius.circular(10),
                border: Border.all(color: Color(0xFF2962FF))),
            child: Container(
              padding: EdgeInsets.all(16),
              child: Text(
                _speechToText.isListening
                    ? '$_lastWords'
                    : _speechToText.isNotListening
                        ? 'Tap the microphone to start listening...'
                        : 'Speech not available',
                style: TextStyle(color: Colors.white),
              ),
            ),
          ),
          Divider(height: 10, color: Colors.transparent),
          Row(
            mainAxisAlignment: MainAxisAlignment.center,
            children: [
              Container(
                height: 80,
                width: 80,
                decoration: BoxDecoration(
                    color: Color(0xFF2962FF),
                    borderRadius: BorderRadius.circular(50)),
                child: IconButton(
                    icon: Icon(
                       _speechToText.isListening ? Icons.mic : Icons.mic_off,
                      size: 35,
                      color: Colors.white,
                    ),
                    onPressed: () async {
                      // listen Speech to text
                      if (_isSpeaking) return;
                      if (await _speechToText.hasPermission &&
                          _speechToText.isNotListening) {
                        _startListening();
                        print("Is listening");
                      } else if (_speechToText.isListening) {
                        _stopListening();
                        print("stop listening");
                      } else {
                        _initSpeechToText();
                        print("init speech");
                      }
                    }),
              ),
              VerticalDivider(
                width: 10,
              ),
              Container(
                height: 50,
                width: 150,
                decoration: BoxDecoration(
                    color: Colors.red, borderRadius: BorderRadius.circular(50)),
                child: ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Color(0xFF2962FF),
                    foregroundColor: Colors.white,
                  ),
                  onPressed: () {
                    Navigator.of(context).push(MaterialPageRoute(
                      builder: (context) => FeedbackPage(conversationData: conversationManager.getHistoryForApi()),
                    ));
                  },
                  child: Text('Finish Interview'),
                ),
              ),
            ],
          ),
          Divider(height: 10, color: Colors.transparent),
        ],
      ),
    );
  }
}

class Message {
  final String content;
  final String role; 
  final DateTime timestamp;

  Message({required this.content, required this.role, required this.timestamp});
}

class Conversation {
  List<Message> messages = [];

  void addMessage(String content, String role) {
    messages
        .add(Message(content: content, role: role, timestamp: DateTime.now()));
  }

  List<Map<String, dynamic>> toApiFormat() {
    return messages
        .map((message) => {
              "role": message.role,
              "content": message.content,
            })
        .toList();
  }
}

class ConversationManager {
  final Conversation conversation = Conversation();
  void addUserMessage(String message) {
    conversation.addMessage(message, "user");
  }
  void addSystemMessage(String message) {
    conversation.addMessage(message, "system");
  }
  List<Map<String, dynamic>> getHistoryForApi() {
    return conversation.toApiFormat();
  }
}
final ConversationManager conversationManager = ConversationManager();

Future<void> sendMessageToGPT(String userMessage) async {
  conversationManager.addUserMessage(userMessage);
  var payload = {
    "model": "gpt-3.5-turbo",
    "messages": conversationManager.getHistoryForApi(),
    "max_tokens": 150,
  };
  var response = await sendToGPTAPI(payload);
  print(response);
  conversationManager.addSystemMessage(response);
}

Future<String> sendToGPTAPI(Map<String, dynamic> payload) async {
  const String apiKey =
      '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'; 
  const String url = 'https://api.openai.com/v1/completions';

  final headers = {
    'Content-Type': 'application/json',
    'Authorization': 'Bearer $apiKey',
  };

  final String body = json.encode(payload);

  try {
    final response =
        await http.post(Uri.parse(url), headers: headers, body: body);
    if (response.statusCode == 200) {
      final Map<String, dynamic> responseData = json.decode(response.body);
      String textResponse = responseData['choices'][0]['text'].trim();
      return textResponse;
    } else {
      throw Exception('Failed to load post');
    }
  } catch (e) {
    print('Error sending message to GPT-3.5 Turbo: $e');
    throw Exception('Error sending message to GPT-3.5 Turbo: $e');
  }
}
