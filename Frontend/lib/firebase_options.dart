import 'package:firebase_core/firebase_core.dart' show FirebaseOptions;
import 'package:flutter/foundation.dart'
    show defaultTargetPlatform, kIsWeb, TargetPlatform;

class DefaultFirebaseOptions {
  static FirebaseOptions get currentPlatform {
    if (kIsWeb) {
      return web;
    }
    switch (defaultTargetPlatform) {
      case TargetPlatform.android:
        return android;
      case TargetPlatform.iOS:
        return ios;
      case TargetPlatform.macOS:
        return macos;
      case TargetPlatform.windows:
        throw UnsupportedError(
          'DefaultFirebaseOptions have not been configured for windows - '
          'you can reconfigure this by running the FlutterFire CLI again.',
        );
      case TargetPlatform.linux:
        throw UnsupportedError(
          'DefaultFirebaseOptions have not been configured for linux - '
          'you can reconfigure this by running the FlutterFire CLI again.',
        );
      default:
        throw UnsupportedError(
          'DefaultFirebaseOptions are not supported for this platform.',
        );
    }
  }

  static const FirebaseOptions web = FirebaseOptions(
    apiKey: 'AIzaSyDmi0z9QBvUdHwPd__TUWxYU4oAEGXccvY',
    appId: '1:813117776614:web:39deb59664df40b7289127',
    messagingSenderId: '813117776614',
    projectId: 'preform-f68e6',
    authDomain: 'preform-f68e6.firebaseapp.com',
    storageBucket: 'preform-f68e6.appspot.com',
    measurementId: 'G-ZS0Z6DY61Z',
  );

  static const FirebaseOptions android = FirebaseOptions(
    apiKey: 'AIzaSyCuLYwERH4GukswbWv9NInHYdlcfoc4Q1M',
    appId: '1:813117776614:android:a3b5caa0ffb0cf29289127',
    messagingSenderId: '813117776614',
    projectId: 'preform-f68e6',
    storageBucket: 'preform-f68e6.appspot.com',
  );

  static const FirebaseOptions ios = FirebaseOptions(
    apiKey: 'AIzaSyARg9g4CfcRs-DnP59AUzjD4a4uTyETEbE',
    appId: '1:813117776614:ios:03abb6ebb07f031d289127',
    messagingSenderId: '813117776614',
    projectId: 'preform-f68e6',
    storageBucket: 'preform-f68e6.appspot.com',
    iosBundleId: 'com.example.preform',
  );

  static const FirebaseOptions macos = FirebaseOptions(
    apiKey: 'AIzaSyARg9g4CfcRs-DnP59AUzjD4a4uTyETEbE',
    appId: '1:813117776614:ios:03abb6ebb07f031d289127',
    messagingSenderId: '813117776614',
    projectId: 'preform-f68e6',
    storageBucket: 'preform-f68e6.appspot.com',
    iosBundleId: 'com.example.preform',
  );
}
