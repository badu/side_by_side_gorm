## side by side gorm

Some tests to check gorm differences after refactoring

Example (test run on 11th of December 2016) of result produced:

| Test name | Allocs | Bytes | Duration  |
| :-------: | -----: | ----: | --------: 
| *OpenTestConnection* | 60 | 6144 | 14.0008ms |
| OpenTestConnection | 59 | 4624 | 8.0005ms |
| Diffs |  :snail: 1 | :snail: 1520 | :snail: 6.0003ms |
| *RunNewMigration* | 8805 | 3056016 | 4.6832679s |
| RunMigration | 8331 | 1275592 | 5.2242988s |
| Diffs |  :snail: 474 | :snail: 1780424 | :zap: 541.0309ms |
| *StringPrimaryKey* | 609 | 24136 | 238.0136ms |
| StringPrimaryKey | 642 | 36272 | 321.0183ms |
| Diffs |  :zap: 33 | :zap: 12136 | :zap: 83.0047ms |
| *SetTable* | 19451 | 982352 | 823.0471ms |
| SetTable | 19159 | 1580656 | 980.0561ms |
| Diffs |  :snail: 292 | :zap: 598304 | :zap: 157.009ms |
| *ExceptionsWithInvalidSql* | 1354 | 76864 | nothing. |
| ExceptionsWithInvalidSql | 2361 | 1102312 | 1.0001ms |
| Diffs |  :zap: 1007 | :zap: 1025448 | :zap: 1.0001ms |
| *HasTable* | 278 | 10760 | 179.0102ms |
| HasTable | 285 | 18400 | 193.0111ms |
| Diffs |  :zap: 7 | :zap: 7640 | :zap: 14.0009ms |
| *TableName* | 186 | 12544 | nothing. |
| TableName | 161 | 22432 | nothing. |
| Diffs |  :snail: 25 | :zap: 9888 | :zzz: |
| *NullValues* | 1481 | 61504 | 309.0177ms |
| NullValues | 1871 | 480720 | 389.0222ms |
| Diffs |  :zap: 390 | :zap: 419216 | :zap: 80.0045ms |
| *NullValuesWithFirstOrCreate* | 1191 | 60024 | 132.0076ms |
| NullValuesWithFirstOrCreate | 967 | 55704 | 158.009ms |
| Diffs |  :snail: 224 | :snail: 4320 | :zap: 26.0014ms |
| *Transaction* | 4217 | 218960 | 80.0046ms |
| Transaction | 4262 | 630336 | 96.0055ms |
| Diffs |  :zap: 45 | :zap: 411376 | :zap: 16.0009ms |
| *Row* | 2385 | 124496 | 201.0115ms |
| Row | 2412 | 150648 | 265.0151ms |
| Diffs |  :zap: 27 | :zap: 26152 | :zap: 64.0036ms |
| *Rows* | 2403 | 125008 | 187.0107ms |
| Rows | 2419 | 147152 | 247.0141ms |
| Diffs |  :zap: 16 | :zap: 22144 | :zap: 60.0034ms |
| *ScanRows* | 2530 | 131208 | 218.0125ms |
| ScanRows | 2534 | 154552 | 248.0142ms |
| Diffs |  :zap: 4 | :zap: 23344 | :zap: 30.0017ms |
| *Scan* | 2747 | 142576 | 183.0104ms |
| Scan | 2932 | 183888 | 233.0134ms |
| Diffs |  :zap: 185 | :zap: 41312 | :zap: 50.003ms |
| *Raw* | 2938 | 154520 | 265.0151ms |
| Raw | 3138 | 194368 | 283.0162ms |
| Diffs |  :zap: 200 | :zap: 39848 | :zap: 18.0011ms |
| *Group* | 170 | 5872 | nothing. |
| Group | 161 | 6240 | nothing. |
| Diffs |  :snail: 9 | :zap: 368 | :zzz: |
| *Joins* | 3928 | 228272 | 96.0055ms |
| Joins | 4195 | 275256 | 98.0056ms |
| Diffs |  :zap: 267 | :zap: 46984 | :zap: 2.0001ms |
| *JoinsWithSelect* | 1225 | 57224 | 68.0039ms |
| JoinsWithSelect | 1369 | 86160 | 97.0055ms |
| Diffs |  :zap: 144 | :zap: 28936 | :zap: 29.0016ms |
| *Having* | 118 | 5776 | 1.0001ms |
| Having | 200 | 13352 | nothing. |
| Diffs |  :zap: 82 | :zap: 7576 | :snail: 1.0001ms |
| *TimeWithZone* | 4050 | 274096 | 186.0106ms |
| TimeWithZone | 3881 | 283056 | 156.009ms |
| Diffs |  :snail: 169 | :zap: 8960 | :snail: 30.0016ms |
| *Hstore* | 27 | 1104 | nothing. |
| Hstore | 31 | 1408 | nothing. |
| Diffs |  :zap: 4 | :zap: 304 | :zzz: |
| *SetAndGet* | 23 | 1184 | 1ms |
| SetAndGet | 27 | 1600 | nothing. |
| Diffs |  :zap: 4 | :zap: 416 | :snail: 1ms |
| *CompatibilityMode* | 747 | 52632 | nothing. |
| CompatibilityMode | 527 | 35320 | nothing. |
| Diffs |  :snail: 220 | :snail: 17312 | :zzz: |
| *OpenExistingDB* | 1170 | 61872 | 72.0041ms |
| OpenExistingDB | 1110 | 67464 | 76.0044ms |
| Diffs |  :snail: 60 | :zap: 5592 | :zap: 4.0003ms |
| *DdlErrors* | 251 | 14392 | nothing. |
| DdlErrors | 581 | 410312 | 1.0001ms |
| Diffs |  :zap: 330 | :zap: 395920 | :zap: 1.0001ms |
| *OpenWithOneParameter* | 20 | 864 | nothing. |
| OpenWithOneParameter | 23 | 976 | nothing. |
| Diffs |  :zap: 3 | :zap: 112 | :zzz: |
| *BelongsTo* | 10602 | 571632 | 658.0376ms |
| BelongsTo | 11523 | 733784 | 818.0468ms |
| Diffs |  :zap: 921 | :zap: 162152 | :zap: 160.0092ms |
| *BelongsToOverrideForeignKey1* | 347 | 16640 | 1ms |
| BelongsToOverrideForeignKey1 | 341 | 20120 | nothing. |
| Diffs |  :snail: 6 | :zap: 3480 | :snail: 1ms |
| *BelongsToOverrideForeignKey2* | 278 | 13800 | nothing. |
| BelongsToOverrideForeignKey2 | 247 | 17528 | nothing. |
| Diffs |  :snail: 31 | :zap: 3728 | :zzz: |
| *HasOne* | 15533 | 842408 | 777.0445ms |
| HasOne | 15691 | 952688 | 681.0389ms |
| Diffs |  :zap: 158 | :zap: 110280 | :snail: 96.0056ms |
| *HasOneOverrideForeignKey1* | 305 | 19992 | 1.0001ms |
| HasOneOverrideForeignKey1 | 273 | 18248 | nothing. |
| Diffs |  :snail: 32 | :snail: 1744 | :snail: 1.0001ms |
| *HasOneOverrideForeignKey2* | 271 | 13624 | nothing. |
| HasOneOverrideForeignKey2 | 246 | 17464 | nothing. |
| Diffs |  :snail: 25 | :zap: 3840 | :zzz: |
| *HasMany* | 11567 | 647400 | 623.0356ms |
| Many | 12088 | 811616 | 700.0401ms |
| Diffs |  :zap: 521 | :zap: 164216 | :zap: 77.0045ms |
| *HasManyOverrideForeignKey1* | 300 | 15008 | 1ms |
| HasManyOverrideForeignKey1 | 268 | 17600 | nothing. |
| Diffs |  :snail: 32 | :zap: 2592 | :snail: 1ms |
| *HasManyOverrideForeignKey2* | 269 | 15040 | nothing. |
| HasManyOverrideForeignKey2 | 244 | 18896 | nothing. |
| Diffs |  :snail: 25 | :zap: 3856 | :zzz: |
| *ManyToMany* | 25215 | 1349536 | 1.9671125s |
| ManyToMany | 27577 | 1716664 | 2.0371166s |
| Diffs |  :zap: 2362 | :zap: 367128 | :zap: 70.0041ms |
| *Related* | 7764 | 406104 | 92.0052ms |
| Related | 7414 | 438680 | 108.0062ms |
| Diffs |  :snail: 350 | :zap: 32576 | :zap: 16.001ms |
| *ForeignKey* | 53 | 4672 | nothing. |
| ForeignKey | 60 | 6896 | nothing. |
| Diffs |  :zap: 7 | :zap: 2224 | :zzz: |
| *LongForeignKey* | 23 | 992 | nothing. |
| LongForeignKey | 26 | 1056 | nothing. |
| Diffs |  :zap: 3 | :zap: 64 | :zzz: |
| *LongForeignKeyWithShortDest* | 23 | 1008 | nothing. |
| LongForeignKeyWithShortDest | 26 | 1072 | nothing. |
| Diffs |  :zap: 3 | :zap: 64 | :zzz: |
| *HasManyChildrenWithOneStruct* | 701 | 30624 | 64.0037ms |
| HasManyChildrenWithOneStruct | 666 | 43544 | 93.0053ms |
| Diffs |  :snail: 35 | :zap: 12920 | :zap: 29.0016ms |
| *RunCallbacks* | 2797 | 132736 | 162.0093ms |
| RunCallbacks | 2773 | 149208 | 207.0118ms |
| Diffs |  :snail: 24 | :zap: 16472 | :zap: 45.0025ms |
| *CallbacksWithErrors* | 5330 | 244472 | 203.0116ms |
| CallbacksWithErrors | 8796 | 4307160 | 241.0138ms |
| Diffs |  :zap: 3466 | :zap: 4062688 | :zap: 38.0022ms |
| *Create* | 2621 | 140264 | 123.0071ms |
| Create | 2115 | 111616 | 164.0093ms |
| Diffs |  :snail: 506 | :snail: 28648 | :zap: 41.0022ms |
| *CreateWithAutoIncrement* | 31 | 1744 | nothing. |
| CreateWithAutoIncrement | 33 | 1792 | nothing. |
| Diffs |  :zap: 2 | :zap: 48 | :zzz: |
| *CreateWithNoGORMPrimayKey* | 274 | 11960 | 101.0058ms |
| CreateWithNoGORMPrimayKey | 279 | 18584 | 115.0065ms |
| Diffs |  :zap: 5 | :zap: 6624 | :zap: 14.0007ms |
| *CreateWithNoStdPrimaryKeyAndDefaultValues* | 1093 | 50088 | 153.0088ms |
| CreateWithNoStdPrimaryKeyAndDefaultValues | 1187 | 75880 | 234.0134ms |
| Diffs |  :zap: 94 | :zap: 25792 | :zap: 81.0046ms |
| *AnonymousScanner* | 1150 | 59488 | 86.0049ms |
| AnonymousScanner | 1099 | 63352 | 106.0061ms |
| Diffs |  :snail: 51 | :zap: 3864 | :zap: 20.0012ms |
| *AnonymousField* | 1653 | 84680 | 103.0059ms |
| AnonymousField | 1621 | 96680 | 92.0053ms |
| Diffs |  :snail: 32 | :zap: 12000 | :snail: 11.0006ms |
| *SelectWithCreate* | 3104 | 150920 | 269.0154ms |
| SelectWithCreate | 3252 | 205544 | 201.0114ms |
| Diffs |  :zap: 148 | :zap: 54624 | :snail: 68.004ms |
| *OmitWithCreate* | 3285 | 167440 | 158.0091ms |
| OmitWithCreate | 3423 | 217128 | 181.0103ms |
| Diffs |  :zap: 138 | :zap: 49688 | :zap: 23.0012ms |
| *CustomizeColumn* | 903 | 42032 | 426.0244ms |
| CustomizeColumn | 861 | 59328 | 546.0312ms |
| Diffs |  :snail: 42 | :zap: 17296 | :zap: 120.0068ms |
| *CustomColumnAndIgnoredFieldClash* | 161 | 13896 | 225.0129ms |
| CustomColumnAndIgnoredFieldClash | 160 | 10488 | 174.01ms |
| Diffs |  :snail: 1 | :snail: 3408 | :snail: 51.0029ms |
| *ManyToManyWithCustomizedColumn* | 1688 | 77248 | 661.0378ms |
| ManyToManyWithCustomizedColumn | 2080 | 138568 | 607.0347ms |
| Diffs |  :zap: 392 | :zap: 61320 | :snail: 54.0031ms |
| *OneToOneWithCustomizedColumn* | 1574 | 74752 | 675.0387ms |
| OneToOneWithCustomizedColumn | 1564 | 98360 | 710.0406ms |
| Diffs |  :snail: 10 | :zap: 23608 | :zap: 35.0019ms |
| *OneToManyWithCustomizedColumn* | 3365 | 167200 | 648.0371ms |
| OneToManyWithCustomizedColumn | 3510 | 217552 | 740.0423ms |
| Diffs |  :zap: 145 | :zap: 50352 | :zap: 92.0052ms |
| *HasOneWithPartialCustomizedColumn* | 2319 | 113600 | 624.0357ms |
| HasOneWithPartialCustomizedColumn | 2446 | 148360 | 733.0419ms |
| Diffs |  :zap: 127 | :zap: 34760 | :zap: 109.0062ms |
| *BelongsToWithPartialCustomizedColumn* | 2550 | 127208 | 592.0339ms |
| BelongsToWithPartialCustomizedColumn | 2691 | 166624 | 658.0376ms |
| Diffs |  :zap: 141 | :zap: 39416 | :zap: 66.0037ms |
| *Delete* | 2292 | 119856 | 176.0101ms |
| Delete | 2179 | 126256 | 213.0122ms |
| Diffs |  :snail: 113 | :zap: 6400 | :zap: 37.0021ms |
| *InlineDelete* | 2316 | 121584 | 326.0186ms |
| InlineDelete | 2303 | 136744 | 294.0168ms |
| Diffs |  :snail: 13 | :zap: 15160 | :snail: 32.0018ms |
| *SoftDelete* | 1042 | 42584 | 300.0171ms |
| SoftDelete | 1287 | 78440 | 407.0233ms |
| Diffs |  :zap: 245 | :zap: 35856 | :zap: 107.0062ms |
| *PrefixColumnNameForEmbeddedStruct* | 435 | 19872 | nothing. |
| PrefixColumnNameForEmbeddedStruct | 426 | 31368 | 1ms |
| Diffs |  :snail: 9 | :zap: 11496 | :zap: 1ms |
| *SaveAndQueryEmbeddedStruct* | 1287 | 52192 | 244.014ms |
| SaveAndQueryEmbeddedStruct | 1368 | 71312 | 262.015ms |
| Diffs |  :zap: 81 | :zap: 19120 | :zap: 18.001ms |
| *CalculateField* | 494 | 25256 | 1ms |
| CalculateField | 449 | 31256 | nothing. |
| Diffs |  :snail: 45 | :zap: 6000 | :snail: 1ms |
| *JoinTable* | 4044 | 200704 | 559.032ms |
| JoinTable | 4324 | 283632 | 641.0367ms |
| Diffs |  :zap: 280 | :zap: 82928 | :zap: 82.0047ms |
| *Indexes* | 8232 | 395152 | 854.0488ms |
| Indexes | 9817 | 3085904 | 944.054ms |
| Diffs |  :zap: 1585 | :zap: 2690752 | :zap: 90.0052ms |
| *AutoMigration* | 1347 | 52824 | 356.0204ms |
| AutoMigration | 1346 | 59968 | 434.0248ms |
| Diffs |  :snail: 1 | :zap: 7144 | :zap: 78.0044ms |
| *MultipleIndexes* | 2436 | 100240 | 957.0547ms |
| MultipleIndexes | 3086 | 939040 | 1.0250587s |
| Diffs |  :zap: 650 | :zap: 838800 | :zap: 68.004ms |
| *ManyToManyWithMultiPrimaryKeys* | 22 | 1040 | nothing. |
| ManyToManyWithMultiPrimaryKeys | 25 | 1136 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :zzz: |
| *ManyToManyWithCustomizedForeignKeys* | 22 | 1056 | nothing. |
| ManyToManyWithCustomizedForeignKeys | 25 | 1152 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :zzz: |
| *ManyToManyWithCustomizedForeignKeys2* | 22 | 1056 | nothing. |
| ManyToManyWithCustomizedForeignKeys2 | 25 | 1152 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :zzz: |
| *PointerFields* | 2048 | 83112 | 530.0303ms |
| PointerFields | 2563 | 156568 | 601.0344ms |
| Diffs |  :zap: 515 | :zap: 73456 | :zap: 71.0041ms |
| *Polymorphic* | 17033 | 868176 | 1.2550718s |
| Polymorphic | 23324 | 1610984 | 1.7601007s |
| Diffs |  :zap: 6291 | :zap: 742808 | :zap: 505.0289ms |
| *NamedPolymorphic* | 11499 | 626368 | 1.2070691s |
| NamedPolymorphic | 16185 | 1130040 | 1.2330705s |
| Diffs |  :zap: 4686 | :zap: 503672 | :zap: 26.0014ms |
| *Preload* | 22329 | 1073920 | 413.0236ms |
| Preload | 22738 | 1308184 | 452.0259ms |
| Diffs |  :zap: 409 | :zap: 234264 | :zap: 39.0023ms |
| *NestedPreload1* | 1760 | 115704 | 514.0294ms |
| NestedPreload1 | 1956 | 124072 | 654.0374ms |
| Diffs |  :zap: 196 | :zap: 8368 | :zap: 140.008ms |
| *NestedPreload2* | 2217 | 93656 | 767.0439ms |
| NestedPreload2 | 2390 | 144920 | 601.0344ms |
| Diffs |  :zap: 173 | :zap: 51264 | :snail: 166.0095ms |
| *NestedPreload3* | 1977 | 90200 | 591.0338ms |
| NestedPreload3 | 2146 | 132256 | 700.04ms |
| Diffs |  :zap: 169 | :zap: 42056 | :zap: 109.0062ms |
| *NestedPreload4* | 1757 | 74256 | 622.0356ms |
| NestedPreload4 | 1915 | 120272 | 603.0345ms |
| Diffs |  :zap: 158 | :zap: 46016 | :snail: 19.0011ms |
| *NestedPreload5* | 2196 | 90728 | 617.0353ms |
| NestedPreload5 | 2375 | 143048 | 750.0429ms |
| Diffs |  :zap: 179 | :zap: 52320 | :zap: 133.0076ms |
| *NestedPreload6* | 3473 | 140232 | 691.0395ms |
| NestedPreload6 | 3688 | 220968 | 842.0482ms |
| Diffs |  :zap: 215 | :zap: 80736 | :zap: 151.0087ms |
| *NestedPreload7* | 3077 | 127976 | 675.0386ms |
| NestedPreload7 | 3292 | 191784 | 823.0471ms |
| Diffs |  :zap: 215 | :zap: 63808 | :zap: 148.0085ms |
| *NestedPreload8* | 2615 | 106296 | 713.0408ms |
| NestedPreload8 | 2806 | 166664 | 677.0387ms |
| Diffs |  :zap: 191 | :zap: 60368 | :snail: 36.0021ms |
| *NestedPreload9* | 5989 | 251648 | 1.0860621s |
| NestedPreload9 | 6274 | 382672 | 1.066061s |
| Diffs |  :zap: 285 | :zap: 131024 | :snail: 20.0011ms |
| *NestedPreload10* | 2110 | 104952 | 913.0522ms |
| NestedPreload10 | 2260 | 132896 | 885.0506ms |
| Diffs |  :zap: 150 | :zap: 27944 | :snail: 28.0016ms |
| *NestedPreload11* | 1809 | 77640 | 670.0384ms |
| NestedPreload11 | 2017 | 122568 | 685.0391ms |
| Diffs |  :zap: 208 | :zap: 44928 | :zap: 15.0007ms |
| *NestedPreload12* | 2488 | 116928 | 773.0443ms |
| NestedPreload12 | 2689 | 158800 | 890.0509ms |
| Diffs |  :zap: 201 | :zap: 41872 | :zap: 117.0066ms |
| *ManyToManyPreloadWithMultiPrimaryKeys* | 23 | 14624 | nothing. |
| ManyToManyPreloadWithMultiPrimaryKeys | 25 | 1152 | nothing. |
| Diffs |  :zap: 2 | :snail: 13472 | :zzz: |
| *ManyToManyPreloadForNestedPointer* | 6505 | 287888 | 909.052ms |
| ManyToManyPreloadForNestedPointer | 8594 | 542616 | 1.0000572s |
| Diffs |  :zap: 2089 | :zap: 254728 | :zap: 91.0052ms |
| *NestedManyToManyPreload* | 4158 | 182368 | 922.0528ms |
| NestedManyToManyPreload | 5375 | 365544 | 1.0930625s |
| Diffs |  :zap: 1217 | :zap: 183176 | :zap: 171.0097ms |
| *NestedManyToManyPreload2* | 2645 | 120336 | 796.0455ms |
| NestedManyToManyPreload2 | 3317 | 215672 | 770.0441ms |
| Diffs |  :zap: 672 | :zap: 95336 | :snail: 26.0014ms |
| *NestedManyToManyPreload3* | 4440 | 191120 | 908.0519ms |
| NestedManyToManyPreload3 | 5400 | 349896 | 1.0090577s |
| Diffs |  :zap: 960 | :zap: 158776 | :zap: 101.0058ms |
| *NestedManyToManyPreload3ForStruct* | 4651 | 198560 | 899.0514ms |
| NestedManyToManyPreload3ForStruct | 5632 | 360736 | 1.0590606s |
| Diffs |  :zap: 981 | :zap: 162176 | :zap: 160.0092ms |
| *NestedManyToManyPreload4* | 3469 | 151104 | 1.2300704s |
| NestedManyToManyPreload4 | 4211 | 293912 | 1.2290703s |
| Diffs |  :zap: 742 | :zap: 142808 | :snail: 1.0001ms |
| *ManyToManyPreloadForPointer* | 4837 | 225104 | 589.0337ms |
| ManyToManyPreloadForPointer | 6614 | 430280 | 735.0421ms |
| Diffs |  :zap: 1777 | :zap: 205176 | :zap: 146.0084ms |
| *NilPointerSlice* | 1848 | 76616 | 685.0391ms |
| NilPointerSlice | 1996 | 120296 | 804.046ms |
| Diffs |  :zap: 148 | :zap: 43680 | :zap: 119.0069ms |
| *NilPointerSlice2* | 1714 | 74848 | 919.0525ms |
| NilPointerSlice2 | 1847 | 125232 | 914.0523ms |
| Diffs |  :zap: 133 | :zap: 50384 | :snail: 5.0002ms |
| *PrefixedPreloadDuplication* | 4025 | 164840 | 1.2140695s |
| PrefixedPreloadDuplication | 4305 | 253408 | 1.329076s |
| Diffs |  :zap: 280 | :zap: 88568 | :zap: 115.0065ms |
| *FirstAndLast* | 4609 | 244056 | 210.012ms |
| FirstAndLast | 3907 | 216456 | 251.0144ms |
| Diffs |  :snail: 702 | :snail: 27600 | :zap: 41.0024ms |
| *FirstAndLastWithNoStdPrimaryKey* | 1550 | 72416 | 156.0089ms |
| FirstAndLastWithNoStdPrimaryKey | 1582 | 96160 | 179.0102ms |
| Diffs |  :zap: 32 | :zap: 23744 | :zap: 23.0013ms |
| *UIntPrimaryKey* | 565 | 28144 | 1.0001ms |
| UIntPrimaryKey | 485 | 28600 | nothing. |
| Diffs |  :snail: 80 | :zap: 456 | :snail: 1.0001ms |
| *StringPrimaryKeyForNumericValueStartingWithZero* | 492 | 20528 | 1ms |
| StringPrimaryKeyForNumericValueStartingWithZero | 933 | 433544 | 8.0005ms |
| Diffs |  :zap: 441 | :zap: 413016 | :zap: 7.0005ms |
| *FindAsSliceOfPointers* | 20562 | 1285240 | 87.005ms |
| FindAsSliceOfPointers | 15619 | 891760 | 111.0063ms |
| Diffs |  :snail: 4943 | :snail: 393480 | :zap: 24.0013ms |
| *SearchWithPlainSQL* | 10383 | 656168 | 250.0143ms |
| SearchWithPlainSQL | 10103 | 653928 | 284.0163ms |
| Diffs |  :snail: 280 | :snail: 2240 | :zap: 34.002ms |
| *SearchWithStruct* | 7623 | 434656 | 248.0142ms |
| SearchWithStruct | 6311 | 352920 | 276.0158ms |
| Diffs |  :snail: 1312 | :snail: 81736 | :zap: 28.0016ms |
| *SearchWithMap* | 6150 | 338328 | 288.0164ms |
| SearchWithMap | 5292 | 309528 | 362.0207ms |
| Diffs |  :snail: 858 | :snail: 28800 | :zap: 74.0043ms |
| *SearchWithEmptyChain* | 4175 | 225568 | 265.0152ms |
| SearchWithEmptyChain | 3985 | 234656 | 234.0134ms |
| Diffs |  :snail: 190 | :zap: 9088 | :snail: 31.0018ms |
| *Select* | 1054 | 55656 | 65.0037ms |
| Select | 1006 | 58160 | 83.0047ms |
| Diffs |  :snail: 48 | :zap: 2504 | :zap: 18.001ms |
| *OrderAndPluck* | 15445 | 957424 | 225.0129ms |
| OrderAndPluck | 12060 | 702360 | 202.0115ms |
| Diffs |  :snail: 3385 | :snail: 255064 | :snail: 23.0014ms |
| *Limit* | 20266 | 1354656 | 340.0195ms |
| Limit | 15894 | 1044632 | 487.0278ms |
| Diffs |  :snail: 4372 | :snail: 310024 | :zap: 147.0083ms |
| *Offset* | 88238 | 5793736 | 1.5950913s |
| Offset | 68935 | 4362008 | 1.8481057s |
| Diffs |  :snail: 19303 | :snail: 1431728 | :zap: 253.0144ms |
| *Or* | 2509 | 152992 | 236.0135ms |
| Or | 2437 | 148792 | 350.02ms |
| Diffs |  :snail: 72 | :snail: 4200 | :zap: 114.0065ms |
| *Count* | 3251 | 177288 | 209.0119ms |
| Count | 3422 | 208616 | 233.0134ms |
| Diffs |  :zap: 171 | :zap: 31328 | :zap: 24.0015ms |
| *Not* | 22100 | 1187592 | 565.0323ms |
| Not | 21565 | 1530224 | 531.0304ms |
| Diffs |  :snail: 535 | :zap: 342632 | :snail: 34.0019ms |
| *FillSmallerStruct* | 912 | 42584 | 65.0037ms |
| FillSmallerStruct | 958 | 56112 | 70.004ms |
| Diffs |  :zap: 46 | :zap: 13528 | :zap: 5.0003ms |
| *FindOrInitialize* | 7206 | 408616 | 96.0055ms |
| FindOrInitialize | 5245 | 276608 | 91.0052ms |
| Diffs |  :snail: 1961 | :snail: 132008 | :snail: 5.0003ms |
| *FindOrCreate* | 12000 | 647168 | 462.0264ms |
| FindOrCreate | 10425 | 1332888 | 562.0322ms |
| Diffs |  :snail: 1575 | :zap: 685720 | :zap: 100.0058ms |
| *SelectWithEscapedFieldName* | 2263 | 117360 | 260.0148ms |
| SelectWithEscapedFieldName | 2053 | 122200 | 289.0166ms |
| Diffs |  :snail: 210 | :zap: 4840 | :zap: 29.0018ms |
| *SelectWithVariables* | 684 | 34160 | 75.0043ms |
| SelectWithVariables | 655 | 39416 | 117.0066ms |
| Diffs |  :snail: 29 | :zap: 5256 | :zap: 42.0023ms |
| *FirstAndLastWithRaw* | 2662 | 135696 | 222.0127ms |
| FirstAndLastWithRaw | 2546 | 147584 | 237.0136ms |
| Diffs |  :snail: 116 | :zap: 11888 | :zap: 15.0009ms |
| *ScannableSlices* | 2696 | 131728 | 91.0052ms |
| ScannableSlices | 768 | 39080 | 83.0048ms |
| Diffs |  :snail: 1928 | :snail: 92648 | :snail: 8.0004ms |
| *Scopes* | 3643 | 204464 | 261.0149ms |
| Scopes | 3459 | 208544 | 323.0185ms |
| Diffs |  :snail: 184 | :zap: 4080 | :zap: 62.0036ms |
| *Update* | 6858 | 324208 | 518.0296ms |
| Update | 6500 | 346368 | 671.0384ms |
| Diffs |  :snail: 358 | :zap: 22160 | :zap: 153.0088ms |
| *UpdateWithNoStdPrimaryKeyAndDefaultValues* | 2992 | 134448 | 501.0287ms |
| UpdateWithNoStdPrimaryKeyAndDefaultValues | 3012 | 169560 | 582.0333ms |
| Diffs |  :zap: 20 | :zap: 35112 | :zap: 81.0046ms |
| *Updates* | 4785 | 217520 | 334.0191ms |
| Updates | 4600 | 241416 | 400.0229ms |
| Diffs |  :snail: 185 | :zap: 23896 | :zap: 66.0038ms |
| *UpdateColumn* | 3260 | 149064 | 359.0205ms |
| UpdateColumn | 2847 | 145448 | 419.024ms |
| Diffs |  :snail: 413 | :snail: 3616 | :zap: 60.0035ms |
| *SelectWithUpdate* | 7175 | 341808 | 349.0199ms |
| SelectWithUpdate | 7261 | 443032 | 321.0184ms |
| Diffs |  :zap: 86 | :zap: 101224 | :snail: 28.0015ms |
| *SelectWithUpdateWithMap* | 7211 | 344744 | 321.0184ms |
| SelectWithUpdateWithMap | 7286 | 442864 | 319.0182ms |
| Diffs |  :zap: 75 | :zap: 98120 | :snail: 2.0002ms |
| *OmitWithUpdate* | 6039 | 292400 | 279.016ms |
| OmitWithUpdate | 6122 | 372056 | 279.0159ms |
| Diffs |  :zap: 83 | :zap: 79656 | :snail: 100ns |
| *OmitWithUpdateWithMap* | 5879 | 287840 | 192.011ms |
| OmitWithUpdateWithMap | 5979 | 366784 | 240.0138ms |
| Diffs |  :zap: 100 | :zap: 78944 | :zap: 48.0028ms |
| *SelectWithUpdateColumn* | 4394 | 213936 | 165.0095ms |
| SelectWithUpdateColumn | 4073 | 240744 | 208.0119ms |
| Diffs |  :snail: 321 | :zap: 26808 | :zap: 43.0024ms |
| *OmitWithUpdateColumn* | 4395 | 214000 | 167.0095ms |
| OmitWithUpdateColumn | 4068 | 240296 | 175.01ms |
| Diffs |  :snail: 327 | :zap: 26296 | :zap: 8.0005ms |
| *UpdateColumnsSkipsAssociations* | 4327 | 207952 | 267.0153ms |
| UpdateColumnsSkipsAssociations | 4046 | 238440 | 233.0133ms |
| Diffs |  :snail: 281 | :zap: 30488 | :snail: 34.002ms |
| *UpdatesWithBlankValues* | 1291 | 62688 | 126.0072ms |
| UpdatesWithBlankValues | 1124 | 58648 | 156.0089ms |
| Diffs |  :snail: 167 | :snail: 4040 | :zap: 30.0017ms |
| *UpdatesTableWithIgnoredValues* | 435 | 16728 | 125.0072ms |
| UpdatesTableWithIgnoredValues | 527 | 27680 | 176.0101ms |
| Diffs |  :zap: 92 | :zap: 10952 | :zap: 51.0029ms |
| *UpdateDecodeVirtualAttributes* | 1045 | 54264 | 127.0072ms |
| UpdateDecodeVirtualAttributes | 929 | 51008 | 156.009ms |
| Diffs |  :snail: 116 | :snail: 3256 | :zap: 29.0018ms |
| *ToDBNameGenerateFriendlyName* | 120 | 5056 | nothing. |
| ToDBNameGenerateFriendlyName | 123 | 5152 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :zzz: |
| *SkipSaveAssociation* | 1323 | 55952 | 484.0276ms |
| SkipSaveAssociation | 1365 | 74680 | 517.0296ms |
| Diffs |  :zap: 42 | :zap: 18728 | :zap: 33.002ms |
| TOTAL (original) | 610319 | 49232632 | 58.6713562s |
| TOTAL (new) | 618763 | 35990056 | 52.8000201s |
| TOTAL (diffs) | 18446744073709543172 | 13242576 | 5.8713361s |