#!/usr/bin/env bash

echo "Reseting test folders."
rm -rf 'test/source/test-folder'
cp -r 'test/source/backup-test-folder/' './test/source/test-folder/'
rm -rf 'test/target/test-folder'
cp -r 'test/target/backup-test-folder/' './test/target/test-folder/'
echo "Done."
