import 'package:flutter/material.dart';
import 'package:preform/pages/feedback_page.dart';
import 'package:preform/pages/interview_page.dart';
import 'package:preform/widgets/ExploreMockInterviewCard.dart';
import 'package:preform/widgets/bottom_navbar.dart';
import 'package:preform/widgets/search_bar.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';

void main(){
  runApp(const MaterialApp(home:ExploreInterviewPage()));
}

class ExploreInterviewPage extends StatefulWidget {
  const ExploreInterviewPage({super.key});

  @override
  State<ExploreInterviewPage> createState() => _ExploreInterviewPageState();
}

class _ExploreInterviewPageState extends State<ExploreInterviewPage> {
  int _selectedIndex = 1;

  void _onNavBarTap(int index) {
    setState(() {
      _selectedIndex = index;
    });

    switch (index) {
      case 0:
        Navigator.pushNamed(context, '/home');
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
      body: SingleChildScrollView(
        child: Padding(
            padding: EdgeInsets.only(left: 20, right: 20, top: 50),
            child: Column(
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                Container(),
                Text("🔍 Explore Mock Interviews",
                    textAlign: TextAlign.start,
                    style:
                        TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
                SizedBox(height: 6),
                Text("Practice with these top industries mock-up interviews!",
                    style:
                        TextStyle(fontSize: 14, fontWeight: FontWeight.normal)),
                SizedBox(height: 12),
                SearchBarInterview(),
                SizedBox(height: 12),
                ExploreMockInterviewCard(),
                SizedBox(height: 12),
                Text("🎨 Personalized Practice",
                    textAlign: TextAlign.start,
                    style:
                        TextStyle(fontSize: 20, fontWeight: FontWeight.bold)),
                SizedBox(height: 12),
                Text(
                    "Haven’t found an interview that’s suitable for you? Customize your own interview requirements!",
                    style:
                        TextStyle(fontSize: 14, fontWeight: FontWeight.normal)),
                SizedBox(height: 12),
                Container(
                  width: 129,
                  child: SizedBox(
                    width: double.infinity, 
                    child: ElevatedButton(
                      onPressed: () {
                        Navigator.push(
                          context,
                          MaterialPageRoute(
                            builder: (context) => const InterviewPage(),
                          ),
                        );
                      },
                      style: ElevatedButton.styleFrom(
                        padding: const EdgeInsets.all(12.0),
                        shape: RoundedRectangleBorder(
                          borderRadius: BorderRadius.circular(60.0),
                        ),
                        backgroundColor: Color(0xFF2962FF),
                        foregroundColor: Color(0xFF2962FF),
                        elevation: 0, 
                        minimumSize: Size(
                          0,
                          0,
                        ), 
                        tapTargetSize: MaterialTapTargetSize.shrinkWrap,
                      ),
                      child: Expanded(
                        child: Row(
                          mainAxisAlignment: MainAxisAlignment.center, 
                          children: [
                            Expanded(
                              child: Text(
                                'Customize now!',
                                style: TextStyle(
                                  fontSize: 12.0,
                                  fontWeight: FontWeight.bold,
                                  color: Colors.white,
                                ),
                                textAlign: TextAlign.center,
                              ),
                            ),
                          ],
                        ),
                      ),
                    ),
                  ),
                ),
                SizedBox(height: 20),
              ],
            )
        ),
      ),
      bottomNavigationBar: BottomNavBar(
        selectedIndex: _selectedIndex,
        onItemSelected: _onNavBarTap,
      ),
    );
  }
}
