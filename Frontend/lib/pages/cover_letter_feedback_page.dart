import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:get/get_core/src/get_main.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';

class CoverLetterFeedbackPage extends StatefulWidget {
  final String feedbackText; // Added to accept feedback text
  const CoverLetterFeedbackPage({Key? key, required this.feedbackText}) : super(key: key);
  @override
  State<CoverLetterFeedbackPage> createState() => _CoverLetterFeedbackPageState();
}

class _CoverLetterFeedbackPageState extends State<CoverLetterFeedbackPage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        leading: IconButton(
          icon: const Icon(
            Icons.arrow_back_ios_new,
            color: Color(0xFF2962FF),
          ),
          onPressed: () => Get.toNamed('/coverLetterPage')
        ),
        title: const Text(
          "Feedback",
          style: TextStyle(fontWeight: FontWeight.bold),
        ),
        backgroundColor: Colors.white,
      ),
      body: SingleChildScrollView(
        child: Padding(
          padding: const EdgeInsets.all(20),
          child: Column(
            mainAxisAlignment: MainAxisAlignment.start,
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const SizedBox(height: 10),
              const Text(
                "Feedback for your Resume :",
                style: TextStyle(fontSize: 18, fontWeight: FontWeight.bold, color: Color(0xFF2962FF)),
              ),
              const SizedBox(height: 20,),
              Container(
                padding: const EdgeInsets.all(20),
                decoration: BoxDecoration(
                  border: Border.all(color: const Color(0xFF2962FF), width: 2),
                  borderRadius: BorderRadius.circular(8),
                ),
                child: Column(
                  children: [
                    Text(
                      widget.feedbackText, 
                      style: const TextStyle(fontSize: 18),
                    ),
                  ],
                )
              )
            ],
          ),
        ),
      ),
    );
  }
}