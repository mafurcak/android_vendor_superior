#!/bin/sh

# Exports

. $ANDROID_BUILD_TOP/vendor/superior/tools/colors

export Changelog=Changelog.txt

if [ -f $Changelog ];
then
	rm -f $Changelog
fi

touch $Changelog

# Print something to build output
echo ${bldppl}"Generating changelog..."${txtrst}

for i in $(seq 15);
do
export After_Date=`date --date="$i days ago" +%m-%d-%Y`
k=$(expr $i - 1)
	export Until_Date=`date --date="$k days ago" +%m-%d-%Y`

	# Line with after --- until was too long for a small ListView
	echo '====================' >> $Changelog;
	echo  "     "$Until_Date    >> $Changelog;
	echo '====================' >> $Changelog;
	# Cycle through every repo to find commits between 2 dates
	repo forall -pc 'git log --pretty=format:"%h  %s  [%cn]" --decorate --after=$After_Date --until=$Until_Date' >> $Changelog
	echo >> $Changelog;
        echo >> $Changelog;
done

sed -i 's/project/   */g' $Changelog

cp $Changelog $OUT/system/etc/
cp $Changelog $OUT/
rm -rf $Changelog