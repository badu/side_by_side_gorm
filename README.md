## side by side gorm

Some tests to check gorm differences after refactoring

Example (test run on 10th of December 2016) of result produced:

| Test name | Allocs | Bytes | Duration  | Dif Allocs | Dif Bytes | Dif Duration |
| :-------: | -----: | ----: | --------: | ---------: | --------: | -----------: |
| original | 58 | 4528 | 1ms | | | |
| diffs | | | | :zap: 58 | :zap: 4528 | :zap: 1ms |
| 0 OpenTestConnection | 0 | 0 | nothing. | | | |
| original | 8808 | 3056752 | 5.2282991s | | | |
| diffs | | | | :zap: 8748 | :zap: 3052048 | :zap: 5.2282991s |
| 1 RunMigration | 60 | 4704 | nothing. | | | |
| original | 610 | 24216 | 289.0166ms | | | |
| diffs | | | | :snail: 7710 | :snail: 1249712 | :snail: 4.2312419s |
| 2 StringPrimaryKey | 8320 | 1273928 | 4.5202585s | | | |
| original | 19459 | 984528 | 941.0538ms | | | |
| diffs | | | | :zap: 18816 | :zap: 948176 | :zap: 628.0359ms |
| 3 SetTable | 643 | 36352 | 313.0179ms | | | |
| original | 1390 | 77712 | 1.0001ms | | | |
| diffs | | | | :snail: 17773 | :snail: 1503424 | :snail: 953.0545ms |
| 4 ExceptionsWithInvalidSql | 19163 | 1581136 | 954.0546ms | | | |
| original | 279 | 10840 | 168.0096ms | | | |
| diffs | | | | :snail: 2052 | :snail: 1090096 | :zap: 167.0096ms |
| 5 HasTable | 2331 | 1100936 | 1ms | | | |
| original | 188 | 12832 | nothing. | | | |
| diffs | | | | :snail: 98 | :snail: 5648 | :snail: 162.0093ms |
| 6 TableName | 286 | 18480 | 162.0093ms | | | |
| original | 1477 | 60768 | 286.0164ms | | | |
| diffs | | | | :zap: 1315 | :zap: 38256 | :zap: 286.0164ms |
| 7 NullValues | 162 | 22512 | nothing. | | | |
| original | 1192 | 60152 | 125.0071ms | | | |
| diffs | | | | :snail: 684 | :snail: 421112 | :snail: 148.0085ms |
| 8 NullValuesWithFirstOrCreate | 1876 | 481264 | 273.0156ms | | | |
| original | 4211 | 217520 | 63.0036ms | | | |
| diffs | | | | :zap: 3243 | :zap: 161736 | :snail: 61.0035ms |
| 9 Transaction | 968 | 55784 | 124.0071ms | | | |
| original | 2386 | 124512 | 187.0107ms | | | |
| diffs | | | | :snail: 1876 | :snail: 505760 | :zap: 125.0071ms |
| 10 Row | 4262 | 630272 | 62.0036ms | | | |
| original | 2406 | 125664 | 162.0093ms | | | |
| diffs | | | | :snail: 11 | :snail: 23800 | :snail: 26.0014ms |
| 11 Rows | 2417 | 149464 | 188.0107ms | | | |
| original | 2530 | 131016 | 222.0127ms | | | |
| diffs | | | | :zap: 109 | :snail: 16488 | :snail: 56.0032ms |
| 12 ScanRows | 2421 | 147504 | 278.0159ms | | | |
| original | 2748 | 142592 | 215.0123ms | | | |
| diffs | | | | :zap: 213 | :snail: 12040 | :snail: 3.0002ms |
| 13 Scan | 2535 | 154632 | 218.0125ms | | | |
| original | 2939 | 154600 | 284.0162ms | | | |
| diffs | | | | :zap: 6 | :snail: 29368 | :zap: 71.004ms |
| 14 Raw | 2933 | 183968 | 213.0122ms | | | |
| original | 171 | 5952 | nothing. | | | |
| diffs | | | | :snail: 2968 | :snail: 188496 | :snail: 332.019ms |
| 15 Group | 3139 | 194448 | 332.019ms | | | |
| original | 3926 | 227984 | 89.0051ms | | | |
| diffs | | | | :zap: 3764 | :zap: 221664 | :zap: 89.0051ms |
| 16 Joins | 162 | 6320 | nothing. | | | |
| original | 1226 | 57304 | 86.0049ms | | | |
| diffs | | | | :snail: 2974 | :snail: 218736 | :snail: 27.0016ms |
| 17 JoinsWithSelect | 4200 | 276040 | 113.0065ms | | | |
| original | 119 | 5856 | nothing. | | | |
| diffs | | | | :snail: 1251 | :snail: 80448 | :snail: 88.005ms |
| 18 Having | 1370 | 86304 | 88.005ms | | | |
| original | 4050 | 274032 | 302.0173ms | | | |
| diffs | | | | :zap: 3849 | :zap: 260600 | :zap: 302.0173ms |
| 19 TimeWithZone | 201 | 13432 | nothing. | | | |
| original | 28 | 1184 | 1.0001ms | | | |
| diffs | | | | :snail: 3853 | :snail: 281744 | :snail: 187.0106ms |
| 20 Hstore | 3881 | 282928 | 188.0107ms | | | |
| original | 24 | 1264 | nothing. | | | |
| diffs | | | | :snail: 7 | :snail: 16 | :zzz: |
| 21 SetAndGet | 31 | 1280 | nothing. | | | |
| original | 749 | 53128 | nothing. | | | |
| diffs | | | | :zap: 720 | :zap: 51240 | :zzz: |
| 22 CompatibilityMode | 29 | 1888 | nothing. | | | |
| original | 1170 | 61744 | 77.0044ms | | | |
| diffs | | | | :zap: 641 | :zap: 25928 | :zap: 77.0044ms |
| 23 OpenExistingDB | 529 | 35816 | nothing. | | | |
| original | 271 | 17112 | 1ms | | | |
| diffs | | | | :snail: 839 | :snail: 50288 | :snail: 78.0046ms |
| 24 DdlErrors | 1110 | 67400 | 79.0046ms | | | |
| original | 21 | 944 | nothing. | | | |
| diffs | | | | :snail: 547 | :snail: 409416 | :zzz: |
| 25 OpenWithOneParameter | 568 | 410360 | nothing. | | | |
| original | 10602 | 571680 | 624.0357ms | | | |
| diffs | | | | :zap: 10578 | :zap: 570624 | :zap: 624.0357ms |
| 26 BelongsTo | 24 | 1056 | nothing. | | | |
| original | 349 | 16928 | nothing. | | | |
| diffs | | | | :snail: 11176 | :snail: 717000 | :snail: 912.0521ms |
| 27 BelongsToOverrideForeignKey1 | 11525 | 733928 | 912.0521ms | | | |
| original | 278 | 13672 | nothing. | | | |
| diffs | | | | :snail: 64 | :snail: 6528 | :zzz: |
| 28 BelongsToOverrideForeignKey2 | 342 | 20200 | nothing. | | | |
| original | 15538 | 843288 | 629.036ms | | | |
| diffs | | | | :zap: 15290 | :zap: 825680 | :zap: 629.036ms |
| 29 HasOne | 248 | 17608 | nothing. | | | |
| original | 307 | 20280 | nothing. | | | |
| diffs | | | | :snail: 15384 | :snail: 932264 | :snail: 662.0378ms |
| 30 HasOneOverrideForeignKey1 | 15691 | 952544 | 662.0378ms | | | |
| original | 273 | 13832 | nothing. | | | |
| diffs | | | | :snail: 1 | :snail: 4496 | :snail: 1.0001ms |
| 31 HasOneOverrideForeignKey2 | 274 | 18328 | 1.0001ms | | | |
| original | 11568 | 648088 | 895.0512ms | | | |
| diffs | | | | :zap: 11321 | :zap: 630544 | :zap: 895.0512ms |
| 32 Many | 247 | 17544 | nothing. | | | |
| original | 302 | 15296 | nothing. | | | |
| diffs | | | | :snail: 11787 | :snail: 796144 | :snail: 768.0439ms |
| 33 HasManyOverrideForeignKey1 | 12089 | 811440 | 768.0439ms | | | |
| original | 269 | 14912 | nothing. | | | |
| diffs | | | | :snail: 1 | :snail: 2976 | :zzz: |
| 34 HasManyOverrideForeignKey2 | 270 | 17888 | nothing. | | | |
| original | 25216 | 1349360 | 1.9511116s | | | |
| diffs | | | | :zap: 24972 | :zap: 1330592 | :zap: 1.9511116s |
| 35 ManyToMany | 244 | 18768 | nothing. | | | |
| original | 7768 | 406584 | 106.0061ms | | | |
| diffs | | | | :snail: 19814 | :snail: 1310080 | :snail: 1.697097s |
| 36 Related | 27582 | 1716664 | 1.8031031s | | | |
| original | 54 | 4752 | nothing. | | | |
| diffs | | | | :snail: 7358 | :snail: 434280 | :snail: 96.0055ms |
| 37 ForeignKey | 7412 | 439032 | 96.0055ms | | | |
| original | 24 | 1072 | nothing. | | | |
| diffs | | | | :snail: 37 | :snail: 5904 | :zzz: |
| 38 LongForeignKey | 61 | 6976 | nothing. | | | |
| original | 24 | 1088 | nothing. | | | |
| diffs | | | | :snail: 3 | :snail: 48 | :zzz: |
| 39 LongForeignKeyWithShortDest | 27 | 1136 | nothing. | | | |
| original | 702 | 30704 | 67.0038ms | | | |
| diffs | | | | :zap: 675 | :zap: 29552 | :zap: 67.0038ms |
| 40 HasManyChildrenWithOneStruct | 27 | 1152 | nothing. | | | |
| original | 2798 | 133280 | 201.0115ms | | | |
| diffs | | | | :zap: 2131 | :zap: 89592 | :zap: 129.0073ms |
| 41 RunCallbacks | 667 | 43688 | 72.0042ms | | | |
| original | 5322 | 242936 | 229.0131ms | | | |
| diffs | | | | :zap: 2545 | :zap: 93120 | :zap: 39.0023ms |
| 42 CallbacksWithErrors | 2777 | 149816 | 190.0108ms | | | |
| original | 2623 | 140616 | 148.0085ms | | | |
| diffs | | | | :snail: 6182 | :snail: 4167456 | :snail: 168.0096ms |
| 43 Create | 8805 | 4308072 | 316.0181ms | | | |
| original | 31 | 1776 | nothing. | | | |
| diffs | | | | :snail: 2084 | :snail: 109760 | :snail: 174.0099ms |
| 44 CreateWithAutoIncrement | 2115 | 111536 | 174.0099ms | | | |
| original | 276 | 12248 | 68.0039ms | | | |
| diffs | | | | :zap: 242 | :zap: 10376 | :zap: 68.0039ms |
| 45 CreateWithNoGORMPrimayKey | 34 | 1872 | nothing. | | | |
| original | 1093 | 49864 | 160.0091ms | | | |
| diffs | | | | :zap: 812 | :zap: 30992 | :zap: 87.0049ms |
| 46 CreateWithNoStdPrimaryKeyAndDefaultValues | 281 | 18872 | 73.0042ms | | | |
| original | 1152 | 59712 | 84.0048ms | | | |
| diffs | | | | :snail: 36 | :snail: 16152 | :snail: 90.0052ms |
| 47 AnonymousScanner | 1188 | 75864 | 174.01ms | | | |
| original | 1654 | 84760 | 88.005ms | | | |
| diffs | | | | :zap: 556 | :zap: 21680 | :snail: 3.0002ms |
| 48 AnonymousField | 1098 | 63080 | 91.0052ms | | | |
| original | 3103 | 150680 | 187.0107ms | | | |
| diffs | | | | :zap: 1482 | :zap: 54128 | :zap: 91.0052ms |
| 49 SelectWithCreate | 1621 | 96552 | 96.0055ms | | | |
| original | 3286 | 167776 | 145.0083ms | | | |
| diffs | | | | :zap: 32 | :snail: 37928 | :snail: 52.003ms |
| 50 OmitWithCreate | 3254 | 205704 | 197.0113ms | | | |
| original | 904 | 42016 | 302.0173ms | | | |
| diffs | | | | :snail: 2520 | :snail: 175192 | :zap: 114.0066ms |
| 51 CustomizeColumn | 3424 | 217208 | 188.0107ms | | | |
| original | 162 | 13976 | 122.007ms | | | |
| diffs | | | | :snail: 700 | :snail: 45432 | :snail: 186.0106ms |
| 52 CustomColumnAndIgnoredFieldClash | 862 | 59408 | 308.0176ms | | | |
| original | 1690 | 77440 | 549.0314ms | | | |
| diffs | | | | :zap: 1529 | :zap: 66872 | :zap: 405.0232ms |
| 53 ManyToManyWithCustomizedColumn | 161 | 10568 | 144.0082ms | | | |
| original | 1575 | 74832 | 648.037ms | | | |
| diffs | | | | :snail: 506 | :snail: 63816 | :snail: 28.0017ms |
| 54 OneToOneWithCustomizedColumn | 2081 | 138648 | 676.0387ms | | | |
| original | 3368 | 167696 | 581.0332ms | | | |
| diffs | | | | :zap: 1804 | :zap: 69464 | :snail: 112.0065ms |
| 55 OneToManyWithCustomizedColumn | 1564 | 98232 | 693.0397ms | | | |
| original | 2319 | 113568 | 541.0309ms | | | |
| diffs | | | | :snail: 1191 | :snail: 103808 | :snail: 126.0073ms |
| 56 HasOneWithPartialCustomizedColumn | 3510 | 217376 | 667.0382ms | | | |
| original | 2551 | 127288 | 545.0312ms | | | |
| diffs | | | | :zap: 103 | :snail: 21264 | :snail: 64.0036ms |
| 57 BelongsToWithPartialCustomizedColumn | 2448 | 148552 | 609.0348ms | | | |
| original | 2294 | 120080 | 200.0115ms | | | |
| diffs | | | | :snail: 398 | :snail: 46624 | :snail: 338.0193ms |
| 58 Delete | 2692 | 166704 | 538.0308ms | | | |
| original | 2317 | 121600 | 270.0154ms | | | |
| diffs | | | | :zap: 137 | :snail: 4640 | :zap: 57.0032ms |
| 59 InlineDelete | 2180 | 126240 | 213.0122ms | | | |
| original | 1043 | 42664 | 262.015ms | | | |
| diffs | | | | :snail: 1261 | :snail: 94224 | :snail: 58.0033ms |
| 60 SoftDelete | 2304 | 136888 | 320.0183ms | | | |
| original | 435 | 19792 | 1.0001ms | | | |
| diffs | | | | :snail: 841 | :snail: 56608 | :snail: 264.015ms |
| 61 PrefixColumnNameForEmbeddedStruct | 1276 | 76400 | 265.0151ms | | | |
| original | 1288 | 52272 | 218.0125ms | | | |
| diffs | | | | :zap: 861 | :zap: 20824 | :zap: 217.0125ms |
| 62 SaveAndQueryEmbeddedStruct | 427 | 31448 | 1ms | | | |
| original | 496 | 25624 | 1.0001ms | | | |
| diffs | | | | :snail: 873 | :snail: 45768 | :snail: 221.0126ms |
| 63 CalculateField | 1369 | 71392 | 222.0127ms | | | |
| original | 4060 | 203160 | 583.0333ms | | | |
| diffs | | | | :zap: 3608 | :zap: 171328 | :zap: 583.0333ms |
| 64 JoinTable | 452 | 31832 | nothing. | | | |
| original | 8222 | 392400 | 804.046ms | | | |
| diffs | | | | :zap: 3896 | :zap: 108400 | :zap: 147.0085ms |
| 65 Indexes | 4326 | 284000 | 657.0375ms | | | |
| original | 1348 | 52904 | 379.0217ms | | | |
| diffs | | | | :snail: 8471 | :snail: 3033832 | :snail: 450.0257ms |
| 66 AutoMigration | 9819 | 3086736 | 829.0474ms | | | |
| original | 2440 | 102240 | 864.0494ms | | | |
| diffs | | | | :zap: 1094 | :zap: 42480 | :zap: 460.0263ms |
| 67 MultipleIndexes | 1346 | 59760 | 404.0231ms | | | |
| original | 23 | 1120 | nothing. | | | |
| diffs | | | | :snail: 3054 | :snail: 936056 | :snail: 849.0486ms |
| 68 ManyToManyWithMultiPrimaryKeys | 3077 | 937176 | 849.0486ms | | | |
| original | 23 | 1136 | nothing. | | | |
| diffs | | | | :snail: 3 | :snail: 80 | :zzz: |
| 69 ManyToManyWithCustomizedForeignKeys | 26 | 1216 | nothing. | | | |
| original | 23 | 1136 | nothing. | | | |
| diffs | | | | :snail: 3 | :snail: 96 | :zzz: |
| 70 ManyToManyWithCustomizedForeignKeys2 | 26 | 1232 | nothing. | | | |
| original | 2049 | 83432 | 444.0254ms | | | |
| diffs | | | | :zap: 2023 | :zap: 82200 | :zap: 444.0254ms |
| 71 PointerFields | 26 | 1232 | nothing. | | | |
| original | 17045 | 870600 | 1.1560661s | | | |
| diffs | | | | :zap: 14481 | :zap: 713744 | :zap: 688.0393ms |
| 72 Polymorphic | 2564 | 156856 | 468.0268ms | | | |
| original | 11500 | 626448 | 1.0710612s | | | |
| diffs | | | | :snail: 11825 | :snail: 984616 | :snail: 199.0115ms |
| 73 NamedPolymorphic | 23325 | 1611064 | 1.2700727s | | | |
| original | 22314 | 1071224 | 445.0255ms | | | |
| diffs | | | | :zap: 6128 | :snail: 58896 | :snail: 809.0463ms |
| 74 Preload | 16186 | 1130120 | 1.2540718s | | | |
| original | 1761 | 116088 | 563.0322ms | | | |
| diffs | | | | :snail: 20999 | :snail: 1196632 | :zap: 141.0081ms |
| 75 NestedPreload1 | 22760 | 1312720 | 422.0241ms | | | |
| original | 2219 | 94056 | 593.0339ms | | | |
| diffs | | | | :zap: 262 | :snail: 30096 | :zap: 39.0022ms |
| 76 NestedPreload2 | 1957 | 124152 | 554.0317ms | | | |
| original | 1978 | 90280 | 557.0319ms | | | |
| diffs | | | | :snail: 412 | :snail: 54336 | :zap: 8.0005ms |
| 77 NestedPreload3 | 2390 | 144616 | 549.0314ms | | | |
| original | 1758 | 74336 | 560.032ms | | | |
| diffs | | | | :snail: 388 | :snail: 57792 | :zap: 9.0005ms |
| 78 NestedPreload4 | 2146 | 132128 | 551.0315ms | | | |
| original | 2194 | 90248 | 612.035ms | | | |
| diffs | | | | :zap: 278 | :snail: 30104 | :zap: 47.0026ms |
| 79 NestedPreload5 | 1916 | 120352 | 565.0324ms | | | |
| original | 3473 | 139960 | 627.0359ms | | | |
| diffs | | | | :zap: 1096 | :snail: 3312 | :snail: 44.0024ms |
| 80 NestedPreload6 | 2377 | 143272 | 671.0383ms | | | |
| original | 3080 | 128616 | 616.0353ms | | | |
| diffs | | | | :snail: 609 | :snail: 92256 | :snail: 82.0046ms |
| 81 NestedPreload7 | 3689 | 220872 | 698.0399ms | | | |
| original | 2614 | 106056 | 662.0378ms | | | |
| diffs | | | | :snail: 677 | :snail: 85344 | :zap: 36.002ms |
| 82 NestedPreload8 | 3291 | 191400 | 626.0358ms | | | |
| original | 5993 | 252480 | 894.0511ms | | | |
| diffs | | | | :zap: 3185 | :zap: 85624 | :zap: 148.0084ms |
| 83 NestedPreload9 | 2808 | 166856 | 746.0427ms | | | |
| original | 2112 | 105432 | 753.0431ms | | | |
| diffs | | | | :snail: 4166 | :snail: 277896 | :snail: 245.014ms |
| 84 NestedPreload10 | 6278 | 383328 | 998.0571ms | | | |
| original | 1809 | 77608 | 581.0332ms | | | |
| diffs | | | | :snail: 454 | :snail: 56024 | :snail: 174.01ms |
| 85 NestedPreload11 | 2263 | 133632 | 755.0432ms | | | |
| original | 2493 | 117792 | 671.0383ms | | | |
| diffs | | | | :zap: 475 | :snail: 4680 | :zap: 95.0053ms |
| 86 NestedPreload12 | 2018 | 122472 | 576.033ms | | | |
| original | 24 | 14704 | nothing. | | | |
| diffs | | | | :snail: 2666 | :snail: 144176 | :snail: 795.0455ms |
| 87 ManyToManyPreloadWithMultiPrimaryKeys | 2690 | 158880 | 795.0455ms | | | |
| original | 6490 | 285432 | 789.0452ms | | | |
| diffs | | | | :zap: 6464 | :zap: 284200 | :zap: 789.0452ms |
| 88 ManyToManyPreloadForNestedPointer | 26 | 1232 | nothing. | | | |
| original | 4158 | 182336 | 949.0543ms | | | |
| diffs | | | | :snail: 4454 | :snail: 363088 | :zap: 44.0025ms |
| 89 NestedManyToManyPreload | 8612 | 545424 | 905.0518ms | | | |
| original | 2646 | 120624 | 851.0487ms | | | |
| diffs | | | | :snail: 2729 | :snail: 244856 | :snail: 214.0122ms |
| 90 NestedManyToManyPreload2 | 5375 | 365480 | 1.0650609s | | | |
| original | 4439 | 190624 | 1.0180582s | | | |
| diffs | | | | :zap: 1121 | :snail: 25128 | :zap: 369.0211ms |
| 91 NestedManyToManyPreload3 | 3318 | 215752 | 649.0371ms | | | |
| original | 4653 | 198784 | 1.0210584s | | | |
| diffs | | | | :snail: 751 | :snail: 151656 | :zap: 81.0046ms |
| 92 NestedManyToManyPreload3ForStruct | 5404 | 350440 | 940.0538ms | | | |
| original | 3469 | 150832 | 968.0554ms | | | |
| diffs | | | | :snail: 2162 | :snail: 209696 | :zap: 22.0013ms |
| 93 NestedManyToManyPreload4 | 5631 | 360528 | 946.0541ms | | | |
| original | 4837 | 225072 | 702.0401ms | | | |
| diffs | | | | :zap: 624 | :snail: 69096 | :snail: 263.0151ms |
| 94 ManyToManyPreloadForPointer | 4213 | 294168 | 965.0552ms | | | |
| original | 1851 | 76920 | 703.0402ms | | | |
| diffs | | | | :snail: 4763 | :snail: 353232 | :snail: 29.0017ms |
| 95 NilPointerSlice | 6614 | 430152 | 732.0419ms | | | |
| original | 1715 | 74928 | 945.054ms | | | |
| diffs | | | | :snail: 284 | :snail: 45672 | :zap: 204.0116ms |
| 96 NilPointerSlice2 | 1999 | 120600 | 741.0424ms | | | |
| original | 4024 | 164520 | 1.1480657s | | | |
| diffs | | | | :zap: 2177 | :zap: 39320 | :zap: 369.0212ms |
| 97 PrefixedPreloadDuplication | 1847 | 125200 | 779.0445ms | | | |
| original | 4612 | 244616 | 216.0124ms | | | |
| diffs | | | | :zap: 306 | :snail: 9048 | :snail: 1.0080576s |
| 98 FirstAndLast | 4306 | 253664 | 1.22407s | | | |
| original | 1550 | 72288 | 161.0092ms | | | |
| diffs | | | | :snail: 2357 | :snail: 144232 | :snail: 35.002ms |
| 99 FirstAndLastWithNoStdPrimaryKey | 3907 | 216520 | 196.0112ms | | | |
| original | 566 | 28192 | nothing. | | | |
| diffs | | | | :snail: 1017 | :snail: 68176 | :snail: 155.0089ms |
| 100 UIntPrimaryKey | 1583 | 96368 | 155.0089ms | | | |
| original | 492 | 20400 | nothing. | | | |
| diffs | | | | :zap: 6 | :snail: 8248 | :zzz: |
| 101 StringPrimaryKeyForNumericValueStartingWithZero | 486 | 28648 | nothing. | | | |
| original | 20582 | 1289072 | 66.0038ms | | | |
| diffs | | | | :zap: 19658 | :zap: 857280 | :zap: 65.0037ms |
| 102 FindAsSliceOfPointers | 924 | 431792 | 1.0001ms | | | |
| original | 10383 | 656136 | 238.0136ms | | | |
| diffs | | | | :snail: 5251 | :snail: 238584 | :zap: 145.0083ms |
| 103 SearchWithPlainSQL | 15634 | 894720 | 93.0053ms | | | |
| original | 7620 | 434096 | 257.0147ms | | | |
| diffs | | | | :snail: 2486 | :snail: 220264 | :zap: 13.0007ms |
| 104 SearchWithStruct | 10106 | 654360 | 244.014ms | | | |
| original | 6151 | 338472 | 266.0152ms | | | |
| diffs | | | | :snail: 163 | :snail: 15072 | :zap: 41.0023ms |
| 105 SearchWithMap | 6314 | 353544 | 225.0129ms | | | |
| original | 4180 | 225768 | 192.011ms | | | |
| diffs | | | | :snail: 1112 | :snail: 83600 | :snail: 83.0047ms |
| 106 SearchWithEmptyChain | 5292 | 309368 | 275.0157ms | | | |
| original | 1055 | 55736 | 66.0038ms | | | |
| diffs | | | | :snail: 2931 | :snail: 179032 | :snail: 158.009ms |
| 107 Select | 3986 | 234768 | 224.0128ms | | | |
| original | 15446 | 957616 | 195.0112ms | | | |
| diffs | | | | :zap: 14439 | :zap: 899376 | :zap: 129.0074ms |
| 108 OrderAndPluck | 1007 | 58240 | 66.0038ms | | | |
| original | 20285 | 1357904 | 385.022ms | | | |
| diffs | | | | :zap: 8224 | :zap: 655464 | :zap: 171.0097ms |
| 109 Limit | 12061 | 702440 | 214.0123ms | | | |
| original | 88257 | 5797400 | 1.5720899s | | | |
| diffs | | | | :zap: 72360 | :zap: 4752624 | :zap: 1.2310704s |
| 110 Offset | 15897 | 1044776 | 341.0195ms | | | |
| original | 2509 | 153024 | 211.012ms | | | |
| diffs | | | | :snail: 66415 | :snail: 4207656 | :snail: 1.379079s |
| 111 Or | 68924 | 4360680 | 1.590091s | | | |
| original | 3249 | 176808 | 200.0115ms | | | |
| diffs | | | | :zap: 811 | :zap: 27872 | :snail: 8.0004ms |
| 112 Count | 2438 | 148936 | 208.0119ms | | | |
| original | 22101 | 1189128 | 514.0294ms | | | |
| diffs | | | | :zap: 18679 | :zap: 980608 | :zap: 339.0194ms |
| 113 Not | 3422 | 208520 | 175.01ms | | | |
| original | 913 | 42664 | 74.0042ms | | | |
| diffs | | | | :snail: 20656 | :snail: 1488264 | :snail: 457.0262ms |
| 114 FillSmallerStruct | 21569 | 1530928 | 531.0304ms | | | |
| original | 7199 | 406672 | 75.0043ms | | | |
| diffs | | | | :zap: 6240 | :zap: 350576 | :snail: 5.0003ms |
| 115 FindOrInitialize | 959 | 56096 | 80.0046ms | | | |
| original | 12007 | 647760 | 455.026ms | | | |
| diffs | | | | :zap: 6753 | :zap: 369528 | :zap: 380.0217ms |
| 116 FindOrCreate | 5254 | 278232 | 75.0043ms | | | |
| original | 2264 | 117376 | 194.0111ms | | | |
| diffs | | | | :snail: 8167 | :snail: 1215432 | :snail: 292.0167ms |
| 117 SelectWithEscapedFieldName | 10431 | 1332808 | 486.0278ms | | | |
| original | 685 | 34240 | 70.004ms | | | |
| diffs | | | | :snail: 1369 | :snail: 88104 | :snail: 160.0091ms |
| 118 SelectWithVariables | 2054 | 122344 | 230.0131ms | | | |
| original | 2664 | 135808 | 173.0099ms | | | |
| diffs | | | | :zap: 2008 | :zap: 96312 | :zap: 93.0053ms |
| 119  fix #1214 : FirstAndLastWithRaw | 656 | 39496 | 80.0046ms | | | |
| original | 2005 | 96720 | 79.0045ms | | | |
| diffs | | | | :snail: 545 | :snail: 51408 | :snail: 124.0071ms |
| 120 ScannableSlices | 2550 | 148128 | 203.0116ms | | | |
| original | 3644 | 204416 | 217.0124ms | | | |
| diffs | | | | :zap: 2134 | :zap: 124376 | :zap: 132.0075ms |
| 121 Scopes | 1510 | 80040 | 85.0049ms | | | |
| original | 6854 | 323744 | 508.029ms | | | |
| diffs | | | | :zap: 3394 | :zap: 115088 | :zap: 292.0166ms |
| 122 Update | 3460 | 208656 | 216.0124ms | | | |
| original | 2993 | 134432 | 436.0249ms | | | |
| diffs | | | | :snail: 3504 | :snail: 211472 | :snail: 38.0023ms |
| 123 UpdateWithNoStdPrimaryKeyAndDefaultValues | 6497 | 345904 | 474.0272ms | | | |
| original | 4789 | 218000 | 310.0177ms | | | |
| diffs | | | | :zap: 1776 | :zap: 48360 | :snail: 236.0135ms |
| 124 Updates | 3013 | 169640 | 546.0312ms | | | |
| original | 3264 | 149576 | 359.0205ms | | | |
| diffs | | | | :snail: 1339 | :snail: 92208 | :zap: 35.0019ms |
| 125 UpdateColumn | 4603 | 241784 | 324.0186ms | | | |
| original | 7175 | 341744 | 333.0191ms | | | |
| diffs | | | | :zap: 4327 | :zap: 196152 | :snail: 24.0013ms |
| 126 SelectWithUpdate | 2848 | 145592 | 357.0204ms | | | |
| original | 7220 | 346304 | 279.016ms | | | |
| diffs | | | | :snail: 32 | :snail: 94608 | :zap: 10.0007ms |
| 127 SelectWithUpdateWithMap | 7252 | 440912 | 269.0153ms | | | |
| original | 6040 | 292480 | 278.0159ms | | | |
| diffs | | | | :snail: 1249 | :snail: 150464 | :snail: 51.0029ms |
| 128 OmitWithUpdate | 7289 | 442944 | 329.0188ms | | | |
| original | 5878 | 287408 | 217.0124ms | | | |
| diffs | | | | :snail: 246 | :snail: 84552 | :snail: 71.0041ms |
| 129 OmitWithUpdateWithMap | 6124 | 371960 | 288.0165ms | | | |
| original | 4393 | 213664 | 166.0095ms | | | |
| diffs | | | | :snail: 1588 | :snail: 153392 | :snail: 67.0038ms |
| 130 SelectWithUpdateColumn | 5981 | 367056 | 233.0133ms | | | |
| original | 4395 | 213808 | 230.0131ms | | | |
| diffs | | | | :zap: 320 | :snail: 27160 | :zap: 45.0025ms |
| 131 OmitWithUpdateColumn | 4075 | 240968 | 185.0106ms | | | |
| original | 4329 | 208240 | 240.0137ms | | | |
| diffs | | | | :zap: 260 | :snail: 31624 | :zap: 63.0036ms |
| 132 UpdateColumnsSkipsAssociations | 4069 | 239864 | 177.0101ms | | | |
| original | 1291 | 62624 | 148.0084ms | | | |
| diffs | | | | :snail: 2753 | :snail: 175336 | :snail: 69.004ms |
| 133 UpdatesWithBlankValues | 4044 | 237960 | 217.0124ms | | | |
| original | 436 | 16808 | 135.0078ms | | | |
| diffs | | | | :snail: 689 | :snail: 41872 | :snail: 5.0002ms |
| 134 UpdatesTableWithIgnoredValues | 1125 | 58680 | 140.008ms | | | |
| original | 1046 | 54312 | 148.0084ms | | | |
| diffs | | | | :zap: 518 | :zap: 26552 | :snail: 4.0003ms |
| 135 UpdateDecodeVirtualAttributes | 528 | 27760 | 152.0087ms | | | |
| original | 122 | 5424 | nothing. | | | |
| diffs | | | | :snail: 808 | :snail: 45664 | :snail: 134.0077ms |
| 136 ToDBNameGenerateFriendlyName | 930 | 51088 | 134.0077ms | | | |
| original | 1324 | 56032 | 448.0257ms | | | |
| diffs | | | | :zap: 1200 | :zap: 50800 | :zap: 448.0257ms |
| 137 SkipSaveAssociation | 124 | 5232 | nothing. | | | |
| TOTAL (original) | 618295 | 35979208 | 51.1929279s | | | |
| TOTAL (new) | 611191 | 49286592 | 52.0219758s | | | |