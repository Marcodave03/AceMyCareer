import 'package:flutter/material.dart';

class CustomExpansionTile extends StatelessWidget {
  final String title;
  final String content;

  const CustomExpansionTile({
    Key? key,
    required this.title,
    required this.content,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Card(
      shape: RoundedRectangleBorder(
        side: const BorderSide(color: Color(0xFF2962FF), width: 2.0),
        borderRadius: BorderRadius.circular(12),
      ),
      color: const Color(0xFF2962FF),
      elevation: 2,
      child: Theme(
        data: Theme.of(context).copyWith(
          dividerColor: const Color(0xFF2962FF),
        ),
        child: ExpansionTile(
          title: Text(
            title,
            style: const TextStyle(
              fontWeight: FontWeight.bold,
              color: Colors.white,
            ),
          ),
          iconColor: Colors.white,
          collapsedIconColor: Colors.white,
          backgroundColor: const Color(
              0xFF2962FF), 
          childrenPadding: EdgeInsets.zero,
          tilePadding: const EdgeInsets.symmetric(horizontal: 16.0),
          children: [
            Container(
              color: Colors.white, 
              padding: const EdgeInsets.all(16.0),
              child: Text(
                content,
                style: const TextStyle(
                  color: Colors.black, 
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
