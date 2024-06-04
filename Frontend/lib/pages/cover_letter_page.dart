import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import '../widgets/bottom_navbar.dart';
import 'dart:io';
import 'package:file_picker/file_picker.dart';
import 'package:http/http.dart' as http;
import 'package:read_pdf_text/read_pdf_text.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

import 'cover_letter_feedback_page.dart';

void main() {
  runApp(const MaterialApp(home: CoverLetterPage()));
}

enum CVOption { curriculumVitae, resume }

class CoverLetterPage extends StatefulWidget {
  const CoverLetterPage({Key? key}) : super(key: key);

  @override
  State<CoverLetterPage> createState() => _CoverLetterPageState();
}

class _CoverLetterPageState extends State<CoverLetterPage> {
  CVOption? _cvOption = CVOption.curriculumVitae;
  File? _cvFile;
  void _pickCV() async {
    FilePickerResult? result = await FilePicker.platform.pickFiles(
      type: FileType.custom,
      allowedExtensions: ['jpg', 'png', 'pdf'],
    );

    if (result != null) {
      setState(() {
        _cvFile = File(result.files.single.path!);
      });
    }
  }

  Future<String> analyzeResumeWithGPT(String resumeText) async {
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
              "You are an expert Resume Analyzer. Provide a compact feedback, and recommendation of this resume: $resumeText, maximum 300 words"
        },
        {"role": "user", "content": ""}
      ],
      "max_tokens": 325,
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

  void _analyzeCV() async {
    if (_cvFile != null) {
      try {
        String resumeText = await ReadPdfText.getPDFtext(_cvFile!.path);

        print('PDF to text result: $resumeText');

        Get.toNamed('/loadingPage');

        String textFeedbackFromGPT = await analyzeResumeWithGPT(resumeText);

        Navigator.of(context).push(MaterialPageRoute(
          builder: (context) =>
              CoverLetterFeedbackPage(feedbackText: textFeedbackFromGPT),
        ));
        Get.back();
        Get.to(
            () => CoverLetterFeedbackPage(feedbackText: textFeedbackFromGPT));
      } catch (e) {
        // if there are error
        print(e);
        ScaffoldMessenger.of(context).showSnackBar(
          const SnackBar(
              content: Text('Failed to analyze CV with GPT, please try again')),
        );
      }
    }
    else {
      ScaffoldMessenger.of(context).showSnackBar(
        const SnackBar(content: Text('Please select a file first')),
      );
    }
  }

  void _clearSelection() {
    setState(() {
      _cvFile = null;
    });
  }

// 2 = Cover Letter Page
  int _selectedIndex = 2;

// for bottom navbar navigation
  void _onNavBarTap(int index) {
    setState(() {
      _selectedIndex = index;
    });

    switch (index) {
      case 0:
        Navigator.pushNamed(context, '/');
        break;
      case 1:
        Navigator.pushNamed(context, '/exploreInterviewPage');
        break;
      case 2:
        Navigator.pushNamed(context, '/coverLetterPage');
        break;
      case 3:
        Navigator.pushNamed(context, '/accountPage');
        break;
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: <Widget>[
          const Flexible(
            child: Padding(
                padding: EdgeInsets.only(left: 25, right: 25),
                child: Text(
                  "Analyze your CV or Resume with AI",
                  style: TextStyle(
                      fontWeight: FontWeight.bold,
                      fontSize: 28,
                      color: Color(0xFF2962FF)),
                )),
          ),
          const Divider(
            height: 10,
            thickness: 1,
            color: Colors.transparent,
          ),
          RadioListTile<CVOption>(
            selectedTileColor: const Color(0xFF2962FF),
            title: const Text(
              'Curriculum Vitae',
              style: TextStyle(color: Color(0xFF2962FF)),
            ),
            value: CVOption.curriculumVitae,
            groupValue: _cvOption,
            onChanged: (CVOption? value) {
              setState(() {
                _cvOption = value;
              });
            },
            activeColor: const Color(0xFF2962FF),
          ),
          RadioListTile<CVOption>(
              selectedTileColor: const Color(0xFF2962FF),
              title: const Text(
                'Resume',
                style: TextStyle(color: Color(0xFF2962FF)),
              ),
              value: CVOption.resume,
              groupValue: _cvOption,
              onChanged: (CVOption? value) {
                setState(() {
                  _cvOption = value;
                });
              },
              activeColor: const Color(0xFF2962FF)),
          const SizedBox(height: 20),
          _cvFile != null
              ? Container(
                  padding: const EdgeInsets.all(40),
                  decoration: BoxDecoration(
                    border:
                        Border.all(color: const Color(0xFF2962FF), width: 2),
                    borderRadius: BorderRadius.circular(8),
                  ),
                  child: Column(
                    children: [
                      const Icon(
                        Icons.file_upload,
                        color: Color(0xFF2962FF),
                        size: 24,
                      ),
                      const SizedBox(
                        height: 10,
                      ),
                      Text('Selected File: ${_cvFile!.path.split('/').last}'),
                      TextButton(
                        onPressed: _clearSelection,
                        child: const Text(
                          'Cancel',
                          style: TextStyle(color: Color(0xFF2962FF)),
                        ),
                      ),
                    ],
                  ),
                )
              : Container(
                  padding: const EdgeInsets.all(40),
                  decoration: BoxDecoration(
                    border:
                        Border.all(color: const Color(0xFF2962FF), width: 2),
                    borderRadius: BorderRadius.circular(8),
                  ),
                  child: Column(
                    children: [
                      const Icon(
                        Icons.file_upload,
                        color: Color(0xFF2962FF),
                        size: 24,
                      ),
                      TextButton(
                        onPressed: _pickCV,
                        child: const Text(
                          'Upload CV .pdf',
                          style: TextStyle(
                            color: Color(0xFF2962FF),
                            fontSize: 16,
                          ),
                        ),
                      ),
                    ],
                  ),
                ),
          const SizedBox(height: 20),
          Padding(
            padding: const EdgeInsets.all(10),
            child: ElevatedButton(
              onPressed: _analyzeCV,
              style: ElevatedButton.styleFrom(
                minimumSize: const Size(double.infinity, 35),
                alignment: Alignment.centerLeft,
                backgroundColor: const Color(0xFF2962FF),
              ),
              child: const Row(
                mainAxisAlignment: MainAxisAlignment.center,
                children: [
                  Text(
                    'Analyze',
                    style: TextStyle(
                        fontWeight: FontWeight.bold,
                        fontSize: 18,
                        color: Colors.white),
                    textAlign: TextAlign.center,
                  ),
                ],
              ),
            ),
          )
        ],
      ),
      bottomNavigationBar: BottomNavBar(
        selectedIndex: _selectedIndex,
        onItemSelected: _onNavBarTap,
      ),
    );
  }
}
