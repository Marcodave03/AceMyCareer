import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;
import 'dart:convert';
import '../widgets/bottom_navbar.dart';

void main() {
  runApp(const MaterialApp(home: AccountPage()));
}

class AccountPage extends StatefulWidget {
  const AccountPage({Key? key}) : super(key: key);

  @override
  State<AccountPage> createState() => _AccountPageState();
}

class Instructor {
  final String firstName;
  final String lastName;
  final String email;
  final String imageUrl; // Add imageUrl field

  Instructor({
    required this.firstName,
    required this.lastName,
    required this.email,
    required this.imageUrl, // Initialize imageUrl
  });

  factory Instructor.fromJson(Map<String, dynamic> json) {
    return Instructor(
      firstName: json['firstname'],
      lastName: json['lastname'],
      email: json['email'],
      imageUrl: json['profile_picture_url'], // Parse imageUrl
    );
  }
}

class _AccountPageState extends State<AccountPage> {
  int _selectedIndex = 3;
  String name = '';
  String email = '';
  String imageUrl = ''; 

  @override
  void initState() {
    super.initState();
    fetchUserData().then((instructor) {
      setState(() {
        name = '${instructor.firstName} ${instructor.lastName}';
        email = instructor.email;
        imageUrl = instructor.imageUrl; // Set imageUrl
      });
    }).catchError((error) {
      print('Error fetching user data: $error');
    });
  }

  Future<Instructor> fetchUserData() async {
    final response = await http.get(Uri.parse('http://localhost:8080/api/users'));

    if (response.statusCode == 200) {
      final List<dynamic> responseData = json.decode(response.body);
      final Map<String, dynamic> userData = responseData.first;
      return Instructor.fromJson(userData);
    } else {
      print('Failed to load user data. Status code: ${response.statusCode}');
      print('Response body: ${response.body}');
      throw Exception('Failed to load user data');
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

  void _editProfile() {
    print('Edit profile tapped');
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.white,
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            CircleAvatar(
              radius: 50,
              backgroundImage: imageUrl.isNotEmpty
                  ? NetworkImage(imageUrl)
                  : AssetImage('./lib/images/chelsea.jpeg') as ImageProvider,
            ),
            SizedBox(height: 20),
            Text(
              name,
              style: TextStyle(
                fontSize: 24,
                fontWeight: FontWeight.bold,
              ),
            ),
            SizedBox(height: 10),
            Text(
              email,
              style: TextStyle(
                fontSize: 18,
                color: Colors.grey,
              ),
            ),
            SizedBox(height: 20),
            ElevatedButton(
              style: ElevatedButton.styleFrom(
                backgroundColor: Color(0xFFC4084F),
                foregroundColor: Colors.white,
              ),
              onPressed: _editProfile,
              child: Text('Edit Profile'),
            ),
          ],
        ),
      ),
      bottomNavigationBar: BottomNavBar(
        selectedIndex: _selectedIndex,
        onItemSelected: _onNavBarTap,
      ),
    );
  }
}
