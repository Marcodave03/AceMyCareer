import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import '../widgets/bottom_navbar.dart';

void main() {
  runApp(const MaterialApp(home: InterviewPage()));
}

class InterviewPage extends StatefulWidget {
  const InterviewPage({Key? key}) : super(key: key);

  @override
  State<InterviewPage> createState() => _InterviewPageState();
}

class InterviewLevel {
  final int id;
  final String name;

  InterviewLevel({
    required this.id,
    required this.name,
  });

  factory InterviewLevel.fromJson(Map<String, dynamic> json) {
    return InterviewLevel(
      id: json['id'],
      name: json['name'],
    );
  }
}

class Position {
  final int id;
  final String name;
  Position({
    required this.id,
    required this.name,
  });

  factory Position.fromJson(Map<String, dynamic> json) {
    return Position(
      id: json['id'],
      name: json['name'],
    );
  }
}

class Industry {
  final int id;
  final String name;
  Industry({
    required this.id,
    required this.name,
  });

  factory Industry.fromJson(Map<String, dynamic> json) {
    return Industry(
      id: json['id'],
      name: json['name'],
    );
  }
}

class _InterviewPageState extends State<InterviewPage> {
  int _selectedIndex = 1;

  List<InterviewLevel> _interviewLevels = [];
  List<bool> levelcheckboxValues = [];
  List<Position> _position = [];
  List<Industry> _industry = [];

  @override
  void initState() {
    super.initState();
    fetchLevelData().then((levels) {
      setState(() {
        _interviewLevels = levels;
        levelcheckboxValues = List<bool>.filled(levels.length, false);
      });
    }).catchError((error) {
      print('Error fetching level data: $error');
    });

    fetchPosition().then((position) {
      setState(() {
        _position = position;
      });
    }).catchError((error) {
      print('Error fetching position data: $error');
    });

    fetchIndustry().then((industry) {
      setState(() {
        _industry = industry;
      });
    }).catchError((error) {
      print('Error fetching industry data: $error');
    });
  }

  Future<List<InterviewLevel>> fetchLevelData() async {
    try {
      final response = await http.get(Uri.parse('http://localhost:8080/api/levels'));
      if (response.statusCode == 200) {
        final List<dynamic> responseData = json.decode(response.body);
        return responseData.map((levelData) => InterviewLevel.fromJson(levelData)).toList();
      } else {
        throw Exception('Failed to load level data. Status code: ${response.statusCode}');
      }
    } catch (error) {
      print('Error fetching level data: $error');
      rethrow;
    }
  }

  Future<List<Position>> fetchPosition() async {
    try {
      final response = await http.get(Uri.parse('http://localhost:8080/api/positions'));
      if (response.statusCode == 200) {
        final List<dynamic> responseData = json.decode(response.body);
        return responseData.map((positionData) => Position.fromJson(positionData)).toList();
      } else {
        throw Exception('Failed to load position data. Status code: ${response.statusCode}');
      }
    } catch (error) {
      print('Error fetching position data: $error');
      rethrow;
    }
  }

  Future<List<Industry>> fetchIndustry() async {
    try {
      final response = await http.get(Uri.parse('http://localhost:8080/api/industries'));
      if (response.statusCode == 200) {
        final List<dynamic> responseData = json.decode(response.body);
        return responseData.map((industryData) => Industry.fromJson(industryData)).toList();
      } else {
        throw Exception('Failed to load industry data. Status code: ${response.statusCode}');
      }
    } catch (error) {
      print('Error fetching industry data: $error');
      rethrow;
    }
  }

  void _onNavBarTap(int index) {
    setState(() {
      _selectedIndex = index;
    });

    switch (index) {
      case 0:
        Navigator.pushNamed(context, '/');
        break;
      case 1:
        Navigator.pushNamed(context, '/interviewPage');
        break;
      case 2:
        Navigator.pushNamed(context, '/coverLetterPage');
        break;
      case 3:
        Navigator.pushNamed(context, '/accountPage');
        break;
    }
  }

  String? _chosenValue;
  String? _industryValue;
  final List<String> items = [
    'React Js',
    'Flutter',
    'Golang',
    'PostgreSQL',
    'Oracle'
  ];
  final List<bool> checkboxValues = [false, false, false, false, false];
  final TextEditingController _textFieldController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: SingleChildScrollView(
        child: Container(
          padding: const EdgeInsets.fromLTRB(28.0, 50.0, 28.0, 0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              const Text(
                "🎨 Personalized Practice",
                style: TextStyle(
                  fontSize: 22,
                  color: Color(0xFF2962FF),
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 20),
              const Text(
                "Industry",
                style: TextStyle(
                  fontSize: 18,
                  color: Colors.black,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 10),
              Container(
                decoration: BoxDecoration(
                  border: Border.all(
                    color: const Color(0xFF2962FF),
                    width: 2.0,
                  ),
                  borderRadius: BorderRadius.circular(8.0),
                ),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    ListView.builder(
                      padding: const EdgeInsets.only(
                        top: 0,
                        bottom: 0,
                      ),
                      shrinkWrap: true,
                      physics: const NeverScrollableScrollPhysics(),
                      itemCount: _industry.length,
                      itemBuilder: (context, index) {
                        final industry = _industry[index];
                        return ListTile(
                          contentPadding: const EdgeInsets.symmetric(
                            horizontal: 14,
                            vertical: 1,
                          ),
                          title: Text(industry.name),
                          dense: true,
                          trailing: Checkbox(
                            value: levelcheckboxValues[index],
                            onChanged: (bool? value) {
                              setState(() {
                                levelcheckboxValues[index] = value ?? false;
                              });
                            },
                          ),
                        );
                      },
                    ),
                  ],
                ),
              ),
              PopupMenuButton<String>(
                child: Container(
                  height: 40,
                  padding: const EdgeInsets.symmetric(horizontal: 16.0),
                  decoration: BoxDecoration(
                    border: Border.all(
                      color: const Color(0xFF2962FF),
                      width: 2.0,
                    ),
                    borderRadius: BorderRadius.circular(8.0),
                  ),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(
                        _industryValue ?? 'Select Industry',
                        style: const TextStyle(color: Colors.black),
                      ),
                      const Icon(Icons.arrow_drop_down, color: Colors.black),
                    ],
                  ),
                ),
                onSelected: (String value) {
                  setState(() {
                    _industryValue = value;
                  });
                },
                itemBuilder: (BuildContext context) {
                  return _industry.isEmpty
                      ? [const PopupMenuItem<String>(
                          value: null,
                          child: Text('Loading...'),
                        )]
                      : _industry.map((industry) => PopupMenuItem<String>(
                          value: industry.name,
                          child: Text(industry.name),
                        )).toList();
                },
              ),
              const SizedBox(height: 10),
              const Text(
                "Position",
                style: TextStyle(
                  fontSize: 18,
                  color: Colors.black,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 10),
              Container(
                decoration: BoxDecoration(
                  border: Border.all(
                    color: const Color(0xFF2962FF),
                    width: 2.0,
                  ),
                  borderRadius: BorderRadius.circular(8.0),
                ),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    ListView.builder(
                      padding: const EdgeInsets.only(
                        top: 0,
                        bottom: 0,
                      ),
                      shrinkWrap: true,
                      physics: const NeverScrollableScrollPhysics(),
                      itemCount: _position.length,
                      itemBuilder: (context, index) {
                        final position = _position[index];
                        return ListTile(
                          contentPadding: const EdgeInsets.symmetric(
                            horizontal: 14,
                            vertical: 1,
                          ),
                          title: Text(position.name),
                          dense: true,
                          trailing: Checkbox(
                            value: levelcheckboxValues[index],
                            onChanged: (bool? value) {
                              setState(() {
                                levelcheckboxValues[index] = value ?? false;
                              });
                            },
                          ),
                        );
                      },
                    ),
                  ],
                ),
              ),
              PopupMenuButton<String>(
                child: Container(
                  height: 40,
                  padding: const EdgeInsets.symmetric(horizontal: 16.0),
                  decoration: BoxDecoration(
                    border: Border.all(
                      color: const Color(0xFF2962FF),
                      width: 2.0,
                    ),
                    borderRadius: BorderRadius.circular(8.0),
                  ),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.spaceBetween,
                    children: [
                      Text(
                        _chosenValue ?? 'Select Position',
                        style: const TextStyle(color: Colors.black),
                      ),
                      const Icon(Icons.arrow_drop_down, color: Colors.black),
                    ],
                  ),
                ),
                onSelected: (String value) {
                  setState(() {
                    _chosenValue = value;
                  });
                },
                itemBuilder: (BuildContext context) {
                  return _position.isEmpty
                      ? [const PopupMenuItem<String>(
                          value: null,
                          child: Text('Loading...'),
                        )]
                      : _position.map((position) => PopupMenuItem<String>(
                          value: position.name,
                          child: Text(position.name),
                        )).toList();
                },
              ),
              const SizedBox(height: 10),
              const Text(
                'Technical Requirements',
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 10),
              Container(
                decoration: BoxDecoration(
                  border: Border.all(
                    color: const Color(0xFF2962FF),
                    width: 2.0,
                  ),
                  borderRadius: BorderRadius.circular(8.0),
                ),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    ListView.builder(
                      padding: const EdgeInsets.only(top: 0, bottom: 0),
                      shrinkWrap: true,
                      physics: const NeverScrollableScrollPhysics(),
                      itemCount: items.length,
                      itemBuilder: (context, index) {
                        return ListTile(
                          contentPadding: const EdgeInsets.symmetric(
                              horizontal: 14, vertical: 1),
                          title: Text(items[index]),
                          dense: true,
                          trailing: Checkbox(
                            value: checkboxValues[index],
                            onChanged: (bool? value) {
                              setState(() {
                                checkboxValues[index] = value ?? false;
                              });
                            },
                          ),
                        );
                      },
                    ),
                    Container(
                      alignment: Alignment.center,
                      child: Padding(
                        padding: const EdgeInsets.only(left: 40, right: 40, bottom: 10),
                        child: ElevatedButton(
                          onPressed: () {},
                          style: ElevatedButton.styleFrom(
                            minimumSize: const Size(double.infinity, 30),
                            backgroundColor: const Color(0xFF2962FF),
                            shape: RoundedRectangleBorder(
                              borderRadius: BorderRadius.circular(10.0),
                            ),
                          ),
                          child: const Row(
                            mainAxisAlignment: MainAxisAlignment.center,
                            children: [
                              Icon(
                                Icons.add,
                                size: 20,
                                color: Colors.white,
                              ),
                              Text(
                                'Add your own requirement',
                                style: TextStyle(
                                    fontWeight: FontWeight.bold,
                                    fontSize: 12,
                                    color: Colors.white),
                                textAlign: TextAlign.center,
                              ),
                            ],
                          ),
                        ),
                      ),
                    )
                  ],
                ),
              ),
              const SizedBox(height: 10),
              const Text(
                'Level',
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 10),
              Container(
                decoration: BoxDecoration(
                  border: Border.all(
                    color: const Color(0xFF2962FF),
                    width: 2.0,
                  ),
                  borderRadius: BorderRadius.circular(8.0),
                ),
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    ListView.builder(
                      padding: const EdgeInsets.only(
                        top: 0,
                        bottom: 0,
                      ),
                      shrinkWrap: true,
                      physics: const NeverScrollableScrollPhysics(),
                      itemCount: _interviewLevels.length,
                      itemBuilder: (context, index) {
                        final interviewLevel = _interviewLevels[index];
                        return ListTile(
                          contentPadding: const EdgeInsets.symmetric(
                            horizontal: 14,
                            vertical: 1,
                          ),
                          title: Text(interviewLevel.name),
                          dense: true,
                          trailing: Checkbox(
                            value: levelcheckboxValues[index],
                            onChanged: (bool? value) {
                              setState(() {
                                levelcheckboxValues[index] = value ?? false;
                              });
                            },
                          ),
                        );
                      },
                    ),
                  ],
                ),
              ),
              const SizedBox(height: 10),
              const Text(
                'Work Experience',
                style: TextStyle(
                  fontSize: 18,
                  fontWeight: FontWeight.bold,
                ),
              ),
              const SizedBox(height: 10),
              TextField(
                controller: _textFieldController,
                decoration: InputDecoration(
                  labelText: 'Add your work experience...',
                  focusedBorder: OutlineInputBorder(
                      borderSide: const BorderSide(
                        color: Color(0xFF2962FF),
                        width: 2.0,
                      ),
                      borderRadius: BorderRadius.circular(8.0)),
                  enabledBorder: OutlineInputBorder(
                      borderSide: const BorderSide(
                          color: Color(0xFF2962FF), width: 2.0),
                      borderRadius: BorderRadius.circular(8.0)),
                ),
              ),
              const SizedBox(height: 10),
              ElevatedButton(
                onPressed: () {
                  Get.toNamed("/recordPage");
                },
                style: ElevatedButton.styleFrom(
                  minimumSize: const Size(double.infinity, 35),
                  alignment: Alignment.centerLeft,
                  backgroundColor: const Color(0xFF2962FF),
                ),
                child: const Row(
                  mainAxisAlignment: MainAxisAlignment.spaceBetween,
                  children: [
                    Text(
                      'Start Now',
                      style: TextStyle(
                          fontWeight: FontWeight.bold,
                          fontSize: 16,
                          color: Colors.white),
                      textAlign: TextAlign.center,
                    ),
                    Icon(
                      Icons.arrow_forward_ios,
                      size: 20,
                      color: Colors.white,
                    ),
                  ],
                ),
              ),
              Divider(height: 10, color: Colors.transparent),
            ],
          ),
        ),
      ),
      bottomNavigationBar: BottomNavBar(
        selectedIndex: _selectedIndex,
        onItemSelected: _onNavBarTap,
      ),
    );
  }

  @override
  void dispose() {
    _textFieldController.dispose();
    super.dispose();
  }
}
