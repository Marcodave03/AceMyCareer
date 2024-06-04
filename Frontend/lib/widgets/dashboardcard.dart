import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';

void main() {
  runApp(const MaterialApp(
      home: DashboardCard(
    progressValue: 0.1,
    numberPercentage: 10,
  )));
}

class DashboardCard extends StatefulWidget {
  final double progressValue;
  final int numberPercentage;

  const DashboardCard({
    Key? key,
    required this.progressValue,
    required this.numberPercentage,
  }) : super(key: key);

  @override
  _DashboardCardState createState() => _DashboardCardState();
}

class _DashboardCardState extends State<DashboardCard> {
  @override
  Widget build(BuildContext context) {
    return Stack(alignment: const Alignment(0.6, 0.6), children: [

      Card(
          margin: const EdgeInsets.all(25),
          elevation: 0,
          shape: const RoundedRectangleBorder(
            side: BorderSide(color: Color(0xFF2962FF)),
            borderRadius: BorderRadius.all(Radius.circular(12)),
          ),
          child: Container(
            padding: const EdgeInsets.all(16.0),
            width: 300,
            child: Column(
              mainAxisSize: MainAxisSize.min,
              children: <Widget>[
                Row(
                  children: <Widget>[
                    SvgPicture.asset("lib/images/data_chart_icon.svg"),
                    const SizedBox(width: 13),
                  ],
                ),
                const Padding(
                  padding: EdgeInsets.only(top: 10.0),
                  child: Align(
                    alignment: Alignment.centerLeft,
                    child: Text(
                      'Data Scientist',
                      style: TextStyle(fontWeight: FontWeight.bold),
                    ),
                  ),
                ),
                const SizedBox(height: 4),
                const Padding(
                  padding: EdgeInsets.only(bottom: 0),
                  child: Align(
                    alignment: Alignment.centerLeft,
                    child: Text(
                      'Entry-Level',
                      style: TextStyle(color: Colors.grey),
                    ),
                  ),
                ),
                const SizedBox(height: 4),
                Padding(
                  padding: const EdgeInsets.only(top: 10, bottom: 10),
                  child: Container(
                    decoration: BoxDecoration(
                      border: Border.all(
                        color: Colors.black, 
                        width: 2.0, 
                      ),
                      borderRadius: BorderRadius.circular(10), 
                    ),
                    child: ClipRRect(
                    borderRadius: BorderRadius.circular(10.0),
                    child: LinearProgressIndicator(
                      value: widget.progressValue,
                      minHeight: 10,
                      backgroundColor: Colors.grey[300],
                      valueColor:
                          const AlwaysStoppedAnimation<Color>(Color(0xFF2962FF)),
                    ),
                  ),
                  )

                ),
                Padding(
                  padding: const EdgeInsets.only(bottom: 16.0),
                  child: Row(
                    mainAxisAlignment: MainAxisAlignment.start,
                    children: [
                      const Text(
                        "Score: ",
                        style: TextStyle(fontWeight: FontWeight.bold),
                      ),
                      Text(widget.numberPercentage.toString(),
                          style: const TextStyle(fontWeight: FontWeight.bold)),
                      const Text("/100",
                          style: TextStyle(fontWeight: FontWeight.bold))
                    ],
                  ),
                ),
                Column(children: <Widget>[
                  Container(
                    width: double.infinity,
                    decoration: BoxDecoration(
                      color: const Color(0xFF2962FF),
                      borderRadius: BorderRadius.circular(
                          4.0), 
                    ),
                    child: TextButton(
                      child: const Text('Try Again',
                          style: TextStyle(
                              fontWeight: FontWeight.bold,
                              color: Colors.white)),
                      onPressed: () {
                        /* ... */
                      },
                    ),
                  ),
                ])
              ],
            ),
          )),
    ]);
  }
}
