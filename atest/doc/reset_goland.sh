#!/bin/bash
echo "Hello World !"

rm -rf ~/Library/Preferences/jetbrains.*
rm -rf ~/Library/Preferences/GoLand2019.3/eval/*.key

sed -i'' -e '/evlsprt/d' ~/Library/Preferences/GoLand2019.3/options/other.xml
sed -i'' -e '/evlsprt/d' ~/Library/Preferences/com.apple.java.util.prefs.plist

echo "END"
