## side by side gorm

Some tests to check gorm differences after refactoring

Example (test run on 10th of December 2016) of result produced:

| Test name | Allocs | Bytes | Duration  |
| :-------: | -----: | ----: | --------: 
| *OpenTestConnection* | 57 | 4448 | 1.0001ms |
| OpenTestConnection | 59 | 4624 | nothing. |
| Diffs |  :zap: 2 | :zap: 176 | :snail: 1.0001ms |
| *RunNewMigration* | 8818 | 3058544 | 5.2362995s |
| RunMigration | 8320 | 1274200 | 4.8202757s |
| Diffs |  :snail: 498 | :snail: 1784344 | :snail: 416.0238ms |
| *StringPrimaryKey* | 609 | 24136 | 276.0158ms |
| StringPrimaryKey | 642 | 36272 | 280.016ms |
| Diffs |  :zap: 33 | :zap: 12136 | :zap: 4.0002ms |
| *SetTable* | 19454 | 981952 | 861.0492ms |
| SetTable | 19156 | 1580032 | 946.0541ms |
| Diffs |  :snail: 298 | :zap: 598080 | :zap: 85.0049ms |
| *ExceptionsWithInvalidSql* | 1385 | 78320 | nothing. |
| ExceptionsWithInvalidSql | 2321 | 1100904 | 1ms |
| Diffs |  :zap: 936 | :zap: 1022584 | :zap: 1ms |
| *HasTable* | 288 | 11128 | 147.0084ms |
| HasTable | 285 | 18400 | 165.0094ms |
| Diffs |  :snail: 3 | :zap: 7272 | :zap: 18.001ms |
| *TableName* | 186 | 12544 | nothing. |
| TableName | 161 | 22432 | nothing. |
| Diffs |  :snail: 25 | :zap: 9888 | :zzz: |
| *NullValues* | 1473 | 61008 | 296.0169ms |
| NullValues | 1883 | 480768 | 363.0208ms |
| Diffs |  :zap: 410 | :zap: 419760 | :zap: 67.0039ms |
| *NullValuesWithFirstOrCreate* | 1191 | 60072 | 144.0082ms |
| NullValuesWithFirstOrCreate | 967 | 55704 | 180.0103ms |
| Diffs |  :snail: 224 | :snail: 4368 | :zap: 36.0021ms |
| *Transaction* | 4210 | 217344 | 264.0151ms |
| Transaction | 4263 | 630336 | 112.0064ms |
| Diffs |  :zap: 53 | :zap: 412992 | :snail: 152.0087ms |
| *Row* | 2385 | 124496 | 236.0135ms |
| Row | 2417 | 149912 | 289.0166ms |
| Diffs |  :zap: 32 | :zap: 25416 | :zap: 53.0031ms |
| *Rows* | 2403 | 125008 | 236.0135ms |
| Rows | 2419 | 147088 | 281.016ms |
| Diffs |  :zap: 16 | :zap: 22080 | :zap: 45.0025ms |
| *ScanRows* | 2529 | 130936 | 236.0135ms |
| ScanRows | 2534 | 154552 | 247.0142ms |
| Diffs |  :zap: 5 | :zap: 23616 | :zap: 11.0007ms |
| *Scan* | 2747 | 142576 | 261.0149ms |
| Scan | 2932 | 183888 | 245.014ms |
| Diffs |  :zap: 185 | :zap: 41312 | :snail: 16.0009ms |
| *Raw* | 2938 | 154520 | 294.0168ms |
| Raw | 3138 | 194368 | 332.019ms |
| Diffs |  :zap: 200 | :zap: 39848 | :zap: 38.0022ms |
| *Group* | 170 | 5872 | nothing. |
| Group | 161 | 6240 | nothing. |
| Diffs |  :snail: 9 | :zap: 368 | :zzz: |
| *Joins* | 3927 | 228320 | 80.0045ms |
| Joins | 4197 | 275416 | 88.0051ms |
| Diffs |  :zap: 270 | :zap: 47096 | :zap: 8.0006ms |
| *JoinsWithSelect* | 1225 | 57224 | 86.0049ms |
| JoinsWithSelect | 1369 | 86096 | 88.005ms |
| Diffs |  :zap: 144 | :zap: 28872 | :zap: 2.0001ms |
| *Having* | 118 | 5776 | nothing. |
| Having | 200 | 13352 | 1.0001ms |
| Diffs |  :zap: 82 | :zap: 7576 | :zap: 1.0001ms |
| *TimeWithZone* | 4048 | 273680 | 135.0077ms |
| TimeWithZone | 3882 | 283200 | 165.0094ms |
| Diffs |  :snail: 166 | :zap: 9520 | :zap: 30.0017ms |
| *Hstore* | 27 | 1104 | nothing. |
| Hstore | 30 | 1200 | 1.0001ms |
| Diffs |  :zap: 3 | :zap: 96 | :zap: 1.0001ms |
| *SetAndGet* | 23 | 1184 | nothing. |
| SetAndGet | 27 | 1600 | nothing. |
| Diffs |  :zap: 4 | :zap: 416 | :zzz: |
| *CompatibilityMode* | 747 | 52632 | nothing. |
| CompatibilityMode | 527 | 35320 | nothing. |
| Diffs |  :snail: 220 | :snail: 17312 | :zzz: |
| *OpenExistingDB* | 1169 | 61664 | 67.0038ms |
| OpenExistingDB | 1109 | 67320 | 81.0047ms |
| Diffs |  :snail: 60 | :zap: 5656 | :zap: 14.0009ms |
| *DdlErrors* | 267 | 15384 | 1ms |
| DdlErrors | 566 | 409528 | nothing. |
| Diffs |  :zap: 299 | :zap: 394144 | :snail: 1ms |
| *OpenWithOneParameter* | 20 | 864 | nothing. |
| OpenWithOneParameter | 24 | 1184 | nothing. |
| Diffs |  :zap: 4 | :zap: 320 | :zzz: |
| *BelongsTo* | 10600 | 571392 | 614.0351ms |
| BelongsTo | 11523 | 733592 | 691.0395ms |
| Diffs |  :zap: 923 | :zap: 162200 | :zap: 77.0044ms |
| *BelongsToOverrideForeignKey1* | 347 | 16640 | nothing. |
| BelongsToOverrideForeignKey1 | 341 | 20120 | 1.0001ms |
| Diffs |  :snail: 6 | :zap: 3480 | :zap: 1.0001ms |
| *BelongsToOverrideForeignKey2* | 278 | 13800 | nothing. |
| BelongsToOverrideForeignKey2 | 247 | 17528 | nothing. |
| Diffs |  :snail: 31 | :zap: 3728 | :zzz: |
| *HasOne* | 15535 | 843000 | 749.0428ms |
| HasOne | 15686 | 952128 | 842.0482ms |
| Diffs |  :zap: 151 | :zap: 109128 | :zap: 93.0054ms |
| *HasOneOverrideForeignKey1* | 306 | 20200 | nothing. |
| HasOneOverrideForeignKey1 | 273 | 18248 | nothing. |
| Diffs |  :snail: 33 | :snail: 1952 | :zzz: |
| *HasOneOverrideForeignKey2* | 270 | 13336 | nothing. |
| HasOneOverrideForeignKey2 | 247 | 17752 | nothing. |
| Diffs |  :snail: 23 | :zap: 4416 | :zzz: |
| *HasMany* | 11567 | 648008 | 671.0384ms |
| Many | 12088 | 811616 | 807.0462ms |
| Diffs |  :zap: 521 | :zap: 163608 | :zap: 136.0078ms |
| *HasManyOverrideForeignKey1* | 299 | 14800 | nothing. |
| HasManyOverrideForeignKey1 | 269 | 17712 | nothing. |
| Diffs |  :snail: 30 | :zap: 2912 | :zzz: |
| *HasManyOverrideForeignKey2* | 268 | 14832 | nothing. |
| HasManyOverrideForeignKey2 | 243 | 18688 | 1.0001ms |
| Diffs |  :snail: 25 | :zap: 3856 | :zap: 1.0001ms |
| *ManyToMany* | 25216 | 1349840 | 1.8871079s |
| ManyToMany | 27580 | 1716600 | 2.2191269s |
| Diffs |  :zap: 2364 | :zap: 366760 | :zap: 332.019ms |
| *Related* | 7765 | 406408 | 92.0053ms |
| Related | 7414 | 439448 | 103.0059ms |
| Diffs |  :snail: 351 | :zap: 33040 | :zap: 11.0006ms |
| *ForeignKey* | 53 | 4672 | nothing. |
| ForeignKey | 60 | 6896 | nothing. |
| Diffs |  :zap: 7 | :zap: 2224 | :zzz: |
| *LongForeignKey* | 23 | 992 | nothing. |
| LongForeignKey | 26 | 1056 | nothing. |
| Diffs |  :zap: 3 | :zap: 64 | :zzz: |
| *LongForeignKeyWithShortDest* | 23 | 1008 | nothing. |
| LongForeignKeyWithShortDest | 26 | 1072 | nothing. |
| Diffs |  :zap: 3 | :zap: 64 | :zzz: |
| *HasManyChildrenWithOneStruct* | 701 | 30624 | 75.0043ms |
| HasManyChildrenWithOneStruct | 666 | 43480 | 89.0051ms |
| Diffs |  :snail: 35 | :zap: 12856 | :zap: 14.0008ms |
| *RunCallbacks* | 2799 | 133504 | 241.0138ms |
| RunCallbacks | 2775 | 149592 | 223.0127ms |
| Diffs |  :snail: 24 | :zap: 16088 | :snail: 18.0011ms |
| *CallbacksWithErrors* | 5331 | 244448 | 220.0126ms |
| CallbacksWithErrors | 8813 | 4310304 | 242.0139ms |
| Diffs |  :zap: 3482 | :zap: 4065856 | :zap: 22.0013ms |
| *Create* | 2620 | 139912 | 140.008ms |
| Create | 2115 | 111408 | 140.008ms |
| Diffs |  :snail: 505 | :snail: 28504 | :zzz: |
| *CreateWithAutoIncrement* | 31 | 1744 | nothing. |
| CreateWithAutoIncrement | 33 | 1792 | nothing. |
| Diffs |  :zap: 2 | :zap: 48 | :zzz: |
| *CreateWithNoGORMPrimayKey* | 274 | 11960 | 68.0039ms |
| CreateWithNoGORMPrimayKey | 279 | 18584 | 65.0037ms |
| Diffs |  :zap: 5 | :zap: 6624 | :snail: 3.0002ms |
| *CreateWithNoStdPrimaryKeyAndDefaultValues* | 1092 | 49880 | 218.0124ms |
| CreateWithNoStdPrimaryKeyAndDefaultValues | 1187 | 75816 | 150.0086ms |
| Diffs |  :zap: 95 | :zap: 25936 | :snail: 68.0038ms |
| *AnonymousScanner* | 1152 | 59840 | 66.0038ms |
| AnonymousScanner | 1097 | 63000 | 67.0038ms |
| Diffs |  :snail: 55 | :zap: 3160 | :zap: 1ms |
| *AnonymousField* | 1655 | 85032 | 73.0042ms |
| AnonymousField | 1621 | 96680 | 89.0051ms |
| Diffs |  :snail: 34 | :zap: 11648 | :zap: 16.0009ms |
| *SelectWithCreate* | 3103 | 150904 | 145.0083ms |
| SelectWithCreate | 3251 | 205400 | 140.008ms |
| Diffs |  :zap: 148 | :zap: 54496 | :snail: 5.0003ms |
| *OmitWithCreate* | 3286 | 167904 | 169.0096ms |
| OmitWithCreate | 3424 | 217368 | 160.0092ms |
| Diffs |  :zap: 138 | :zap: 49464 | :snail: 9.0004ms |
| *CustomizeColumn* | 903 | 42032 | 338.0193ms |
| CustomizeColumn | 861 | 59328 | 456.0261ms |
| Diffs |  :snail: 42 | :zap: 17296 | :zap: 118.0068ms |
| *CustomColumnAndIgnoredFieldClash* | 161 | 13896 | 214.0122ms |
| CustomColumnAndIgnoredFieldClash | 160 | 10488 | 175.01ms |
| Diffs |  :snail: 1 | :snail: 3408 | :snail: 39.0022ms |
| *ManyToManyWithCustomizedColumn* | 1687 | 77040 | 675.0386ms |
| ManyToManyWithCustomizedColumn | 2081 | 138856 | 676.0387ms |
| Diffs |  :zap: 394 | :zap: 61816 | :zap: 1.0001ms |
| *OneToOneWithCustomizedColumn* | 1574 | 74752 | 767.0439ms |
| OneToOneWithCustomizedColumn | 1563 | 98248 | 684.0391ms |
| Diffs |  :snail: 11 | :zap: 23496 | :snail: 83.0048ms |
| *OneToManyWithCustomizedColumn* | 3366 | 167056 | 567.0324ms |
| OneToManyWithCustomizedColumn | 3510 | 217280 | 713.0408ms |
| Diffs |  :zap: 144 | :zap: 50224 | :zap: 146.0084ms |
| *HasOneWithPartialCustomizedColumn* | 2318 | 113488 | 602.0344ms |
| HasOneWithPartialCustomizedColumn | 2446 | 148360 | 710.0406ms |
| Diffs |  :zap: 128 | :zap: 34872 | :zap: 108.0062ms |
| *BelongsToWithPartialCustomizedColumn* | 2550 | 127208 | 671.0384ms |
| BelongsToWithPartialCustomizedColumn | 2691 | 166624 | 667.0382ms |
| Diffs |  :zap: 141 | :zap: 39416 | :snail: 4.0002ms |
| *Delete* | 2292 | 119856 | 234.0134ms |
| Delete | 2179 | 126256 | 269.0153ms |
| Diffs |  :snail: 113 | :zap: 6400 | :zap: 35.0019ms |
| *InlineDelete* | 2316 | 121584 | 259.0149ms |
| InlineDelete | 2304 | 137016 | 349.0199ms |
| Diffs |  :snail: 12 | :zap: 15432 | :zap: 90.005ms |
| *SoftDelete* | 1042 | 42584 | 245.014ms |
| SoftDelete | 1276 | 76480 | 296.017ms |
| Diffs |  :zap: 234 | :zap: 33896 | :zap: 51.003ms |
| *PrefixColumnNameForEmbeddedStruct* | 434 | 19712 | 1ms |
| PrefixColumnNameForEmbeddedStruct | 424 | 30952 | 1.0001ms |
| Diffs |  :snail: 10 | :zap: 11240 | :zap: 100ns |
| *SaveAndQueryEmbeddedStruct* | 1288 | 52304 | 255.0146ms |
| SaveAndQueryEmbeddedStruct | 1368 | 71312 | 278.0159ms |
| Diffs |  :zap: 80 | :zap: 19008 | :zap: 23.0013ms |
| *CalculateField* | 494 | 25256 | nothing. |
| CalculateField | 450 | 31464 | nothing. |
| Diffs |  :snail: 44 | :zap: 6208 | :zzz: |
| *JoinTable* | 4061 | 204056 | 534.0306ms |
| JoinTable | 4325 | 283920 | 696.0398ms |
| Diffs |  :zap: 264 | :zap: 79864 | :zap: 162.0092ms |
| *Indexes* | 8231 | 393184 | 973.0557ms |
| Indexes | 9816 | 3085760 | 1.0950626s |
| Diffs |  :zap: 1585 | :zap: 2692576 | :zap: 122.0069ms |
| *AutoMigration* | 1347 | 52824 | 378.0217ms |
| AutoMigration | 1345 | 59680 | 422.0241ms |
| Diffs |  :snail: 2 | :zap: 6856 | :zap: 44.0024ms |
| *MultipleIndexes* | 2438 | 101744 | 1.0630608s |
| MultipleIndexes | 3076 | 937096 | 911.0521ms |
| Diffs |  :zap: 638 | :zap: 835352 | :snail: 152.0087ms |
| *ManyToManyWithMultiPrimaryKeys* | 22 | 1040 | nothing. |
| ManyToManyWithMultiPrimaryKeys | 25 | 1136 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :zzz: |
| *ManyToManyWithCustomizedForeignKeys* | 22 | 1056 | nothing. |
| ManyToManyWithCustomizedForeignKeys | 25 | 1152 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :zzz: |
| *ManyToManyWithCustomizedForeignKeys2* | 22 | 1056 | nothing. |
| ManyToManyWithCustomizedForeignKeys2 | 25 | 1152 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :zzz: |
| *PointerFields* | 2046 | 82856 | 474.0271ms |
| PointerFields | 2569 | 158256 | 626.0358ms |
| Diffs |  :zap: 523 | :zap: 75400 | :zap: 152.0087ms |
| *Polymorphic* | 17034 | 868368 | 1.3600778s |
| Polymorphic | 23324 | 1610984 | 1.5500887s |
| Diffs |  :zap: 6290 | :zap: 742616 | :zap: 190.0109ms |
| *NamedPolymorphic* | 11500 | 626528 | 1.1600663s |
| NamedPolymorphic | 16185 | 1130040 | 1.2150695s |
| Diffs |  :zap: 4685 | :zap: 503512 | :zap: 55.0032ms |
| *Preload* | 22330 | 1073840 | 471.027ms |
| Preload | 22739 | 1309096 | 478.0273ms |
| Diffs |  :zap: 409 | :zap: 235256 | :zap: 7.0003ms |
| *NestedPreload1* | 1759 | 115592 | 645.0369ms |
| NestedPreload1 | 1957 | 124216 | 716.0409ms |
| Diffs |  :zap: 198 | :zap: 8624 | :zap: 71.004ms |
| *NestedPreload2* | 2216 | 93544 | 563.0322ms |
| NestedPreload2 | 2391 | 144952 | 578.0331ms |
| Diffs |  :zap: 175 | :zap: 51408 | :zap: 15.0009ms |
| *NestedPreload3* | 1977 | 90200 | 572.0327ms |
| NestedPreload3 | 2145 | 132048 | 628.0359ms |
| Diffs |  :zap: 168 | :zap: 41848 | :zap: 56.0032ms |
| *NestedPreload4* | 1757 | 74256 | 588.0336ms |
| NestedPreload4 | 1915 | 120272 | 662.0379ms |
| Diffs |  :zap: 158 | :zap: 46016 | :zap: 74.0043ms |
| *NestedPreload5* | 2194 | 90376 | 756.0432ms |
| NestedPreload5 | 2374 | 142936 | 746.0427ms |
| Diffs |  :zap: 180 | :zap: 52560 | :snail: 10.0005ms |
| *NestedPreload6* | 3476 | 140696 | 630.0361ms |
| NestedPreload6 | 3689 | 221112 | 750.0429ms |
| Diffs |  :zap: 213 | :zap: 80416 | :zap: 120.0068ms |
| *NestedPreload7* | 3077 | 127976 | 630.036ms |
| NestedPreload7 | 3290 | 191320 | 703.0403ms |
| Diffs |  :zap: 213 | :zap: 63344 | :zap: 73.0043ms |
| *NestedPreload8* | 2616 | 106536 | 748.0428ms |
| NestedPreload8 | 2806 | 166664 | 783.0448ms |
| Diffs |  :zap: 190 | :zap: 60128 | :zap: 35.002ms |
| *NestedPreload9* | 5990 | 251584 | 1.0230585s |
| NestedPreload9 | 6277 | 383184 | 1.0630608s |
| Diffs |  :zap: 287 | :zap: 131600 | :zap: 40.0023ms |
| *NestedPreload10* | 2110 | 104920 | 737.0422ms |
| NestedPreload10 | 2262 | 133712 | 802.0458ms |
| Diffs |  :zap: 152 | :zap: 28792 | :zap: 65.0036ms |
| *NestedPreload11* | 1807 | 77416 | 661.0379ms |
| NestedPreload11 | 2017 | 122568 | 711.0407ms |
| Diffs |  :zap: 210 | :zap: 45152 | :zap: 50.0028ms |
| *NestedPreload12* | 2492 | 117616 | 734.042ms |
| NestedPreload12 | 2690 | 158944 | 934.0534ms |
| Diffs |  :zap: 198 | :zap: 41328 | :zap: 200.0114ms |
| *ManyToManyPreloadWithMultiPrimaryKeys* | 23 | 14624 | nothing. |
| ManyToManyPreloadWithMultiPrimaryKeys | 25 | 1152 | nothing. |
| Diffs |  :zap: 2 | :snail: 13472 | :zzz: |
| *ManyToManyPreloadForNestedPointer* | 6489 | 285352 | 789.0451ms |
| ManyToManyPreloadForNestedPointer | 8599 | 544240 | 890.0509ms |
| Diffs |  :zap: 2110 | :zap: 258888 | :zap: 101.0058ms |
| *NestedManyToManyPreload* | 4157 | 182256 | 961.055ms |
| NestedManyToManyPreload | 5376 | 365752 | 1.0520602s |
| Diffs |  :zap: 1219 | :zap: 183496 | :zap: 91.0052ms |
| *NestedManyToManyPreload2* | 2647 | 120768 | 835.0477ms |
| NestedManyToManyPreload2 | 3319 | 216024 | 798.0457ms |
| Diffs |  :zap: 672 | :zap: 95256 | :snail: 37.002ms |
| *NestedManyToManyPreload3* | 4440 | 191072 | 1.0050575s |
| NestedManyToManyPreload3 | 5401 | 350040 | 928.0531ms |
| Diffs |  :zap: 961 | :zap: 158968 | :snail: 77.0044ms |
| *NestedManyToManyPreload3ForStruct* | 4653 | 198992 | 938.0537ms |
| NestedManyToManyPreload3ForStruct | 5628 | 360048 | 986.0564ms |
| Diffs |  :zap: 975 | :zap: 161056 | :zap: 48.0027ms |
| *NestedManyToManyPreload4* | 3469 | 150896 | 1.2140694s |
| NestedManyToManyPreload4 | 4210 | 293800 | 1.153066s |
| Diffs |  :zap: 741 | :zap: 142904 | :snail: 61.0034ms |
| *ManyToManyPreloadForPointer* | 4836 | 224928 | 795.0454ms |
| ManyToManyPreloadForPointer | 6613 | 430072 | 697.0399ms |
| Diffs |  :zap: 1777 | :zap: 205144 | :snail: 98.0055ms |
| *NilPointerSlice* | 1848 | 76616 | 701.0401ms |
| NilPointerSlice | 1997 | 120504 | 747.0427ms |
| Diffs |  :zap: 149 | :zap: 43888 | :zap: 46.0026ms |
| *NilPointerSlice2* | 1714 | 74848 | 945.0541ms |
| NilPointerSlice2 | 1846 | 125120 | 1.1060632s |
| Diffs |  :zap: 132 | :zap: 50272 | :zap: 161.0091ms |
| *PrefixedPreloadDuplication* | 4024 | 164552 | 1.3750787s |
| PrefixedPreloadDuplication | 4304 | 253296 | 1.2820733s |
| Diffs |  :zap: 280 | :zap: 88744 | :snail: 93.0054ms |
| *FirstAndLast* | 4609 | 243800 | 205.0117ms |
| FirstAndLast | 3908 | 216728 | 211.012ms |
| Diffs |  :snail: 701 | :snail: 27072 | :zap: 6.0003ms |
| *FirstAndLastWithNoStdPrimaryKey* | 1549 | 72208 | 127.0073ms |
| FirstAndLastWithNoStdPrimaryKey | 1582 | 96224 | 230.0132ms |
| Diffs |  :zap: 33 | :zap: 24016 | :zap: 103.0059ms |
| *UIntPrimaryKey* | 565 | 28144 | 1ms |
| UIntPrimaryKey | 485 | 28568 | nothing. |
| Diffs |  :snail: 80 | :zap: 424 | :snail: 1ms |
| *StringPrimaryKeyForNumericValueStartingWithZero* | 493 | 21152 | 1ms |
| StringPrimaryKeyForNumericValueStartingWithZero | 912 | 431088 | 2.0001ms |
| Diffs |  :zap: 419 | :zap: 409936 | :zap: 1.0001ms |
| *FindAsSliceOfPointers* | 20569 | 1286416 | 75.0043ms |
| FindAsSliceOfPointers | 15633 | 894768 | 73.0042ms |
| Diffs |  :snail: 4936 | :snail: 391648 | :snail: 2.0001ms |
| *SearchWithPlainSQL* | 10386 | 656952 | 265.0152ms |
| SearchWithPlainSQL | 10104 | 653944 | 243.0139ms |
| Diffs |  :snail: 282 | :snail: 3008 | :snail: 22.0013ms |
| *SearchWithStruct* | 7617 | 433600 | 248.0142ms |
| SearchWithStruct | 6311 | 352984 | 267.0152ms |
| Diffs |  :snail: 1306 | :snail: 80616 | :zap: 19.001ms |
| *SearchWithMap* | 6151 | 338440 | 336.0192ms |
| SearchWithMap | 5293 | 309560 | 336.0192ms |
| Diffs |  :snail: 858 | :snail: 28880 | :zzz: |
| *SearchWithEmptyChain* | 4177 | 225792 | 332.019ms |
| SearchWithEmptyChain | 3981 | 233984 | 283.0162ms |
| Diffs |  :snail: 196 | :zap: 8192 | :snail: 49.0028ms |
| *Select* | 1054 | 55656 | 74.0042ms |
| Select | 1006 | 58160 | 83.0048ms |
| Diffs |  :snail: 48 | :zap: 2504 | :zap: 9.0006ms |
| *OrderAndPluck* | 15445 | 957536 | 229.0131ms |
| OrderAndPluck | 12059 | 702152 | 274.0156ms |
| Diffs |  :snail: 3386 | :snail: 255384 | :zap: 45.0025ms |
| *Limit* | 20276 | 1356128 | 408.0234ms |
| Limit | 15894 | 1043992 | 491.0281ms |
| Diffs |  :snail: 4382 | :snail: 312136 | :zap: 83.0047ms |
| *Offset* | 88240 | 5792952 | 1.5800904s |
| Offset | 68937 | 4361224 | 1.9381108s |
| Diffs |  :snail: 19303 | :snail: 1431728 | :zap: 358.0204ms |
| *Or* | 2509 | 153152 | 267.0153ms |
| Or | 2436 | 148648 | 270.0155ms |
| Diffs |  :snail: 73 | :snail: 4504 | :zap: 3.0002ms |
| *Count* | 3248 | 176728 | 267.0152ms |
| Count | 3422 | 208616 | 283.0162ms |
| Diffs |  :zap: 174 | :zap: 31888 | :zap: 16.001ms |
| *Not* | 22105 | 1188952 | 614.0351ms |
| Not | 21568 | 1530720 | 488.028ms |
| Diffs |  :snail: 537 | :zap: 341768 | :snail: 126.0071ms |
| *FillSmallerStruct* | 912 | 42584 | 75.0043ms |
| FillSmallerStruct | 958 | 56112 | 108.0062ms |
| Diffs |  :zap: 46 | :zap: 13528 | :zap: 33.0019ms |
| *FindOrInitialize* | 7196 | 406144 | 57.0032ms |
| FindOrInitialize | 5252 | 278344 | 101.0058ms |
| Diffs |  :snail: 1944 | :snail: 127800 | :zap: 44.0026ms |
| *FindOrCreate* | 12007 | 647328 | 479.0274ms |
| FindOrCreate | 10421 | 1332376 | 551.0315ms |
| Diffs |  :snail: 1586 | :zap: 685048 | :zap: 72.0041ms |
| *SelectWithEscapedFieldName* | 2263 | 117296 | 205.0117ms |
| SelectWithEscapedFieldName | 2053 | 122200 | 247.0141ms |
| Diffs |  :snail: 210 | :zap: 4904 | :zap: 42.0024ms |
| *SelectWithVariables* | 684 | 34160 | 69.004ms |
| SelectWithVariables | 655 | 39416 | 88.0051ms |
| Diffs |  :snail: 29 | :zap: 5256 | :zap: 19.0011ms |
| *FirstAndLastWithRaw* | 2664 | 136112 | 240.0137ms |
| FirstAndLastWithRaw | 2549 | 148080 | 186.0106ms |
| Diffs |  :snail: 115 | :zap: 11968 | :snail: 54.0031ms |
| *ScannableSlices* | 2634 | 128800 | 58.0033ms |
| ScannableSlices | 711 | 35944 | 103.0059ms |
| Diffs |  :snail: 1923 | :snail: 92856 | :zap: 45.0026ms |
| *Scopes* | 3645 | 204688 | 225.0129ms |
| Scopes | 3461 | 208928 | 280.016ms |
| Diffs |  :snail: 184 | :zap: 4240 | :zap: 55.0031ms |
| *Update* | 6850 | 322768 | 607.0347ms |
| Update | 6500 | 346352 | 674.0385ms |
| Diffs |  :snail: 350 | :zap: 23584 | :zap: 67.0038ms |
| *UpdateWithNoStdPrimaryKeyAndDefaultValues* | 2992 | 134272 | 576.033ms |
| UpdateWithNoStdPrimaryKeyAndDefaultValues | 3012 | 169384 | 615.0352ms |
| Diffs |  :zap: 20 | :zap: 35112 | :zap: 39.0022ms |
| *Updates* | 4788 | 217952 | 367.021ms |
| Updates | 4599 | 241272 | 399.0228ms |
| Diffs |  :snail: 189 | :zap: 23320 | :zap: 32.0018ms |
| *UpdateColumn* | 3261 | 149176 | 335.0192ms |
| UpdateColumn | 2848 | 145784 | 399.0228ms |
| Diffs |  :snail: 413 | :snail: 3392 | :zap: 64.0036ms |
| *SelectWithUpdate* | 7175 | 341232 | 274.0157ms |
| SelectWithUpdate | 7259 | 442696 | 287.0164ms |
| Diffs |  :zap: 84 | :zap: 101464 | :zap: 13.0007ms |
| *SelectWithUpdateWithMap* | 7213 | 345432 | 262.015ms |
| SelectWithUpdateWithMap | 7287 | 442400 | 337.0193ms |
| Diffs |  :zap: 74 | :zap: 96968 | :zap: 75.0043ms |
| *OmitWithUpdate* | 6041 | 292256 | 278.0159ms |
| OmitWithUpdate | 6124 | 372776 | 306.0175ms |
| Diffs |  :zap: 83 | :zap: 80520 | :zap: 28.0016ms |
| *OmitWithUpdateWithMap* | 5878 | 287248 | 204.0117ms |
| OmitWithUpdateWithMap | 5979 | 366896 | 237.0135ms |
| Diffs |  :zap: 101 | :zap: 79648 | :zap: 33.0018ms |
| *SelectWithUpdateColumn* | 4394 | 213872 | 165.0095ms |
| SelectWithUpdateColumn | 4072 | 240536 | 224.0128ms |
| Diffs |  :snail: 322 | :zap: 26664 | :zap: 59.0033ms |
| *OmitWithUpdateColumn* | 4393 | 213584 | 184.0105ms |
| OmitWithUpdateColumn | 4070 | 240200 | 191.011ms |
| Diffs |  :snail: 323 | :zap: 26616 | :zap: 7.0005ms |
| *UpdateColumnsSkipsAssociations* | 4329 | 208368 | 268.0154ms |
| UpdateColumnsSkipsAssociations | 4045 | 238232 | 281.016ms |
| Diffs |  :snail: 284 | :zap: 29864 | :zap: 13.0006ms |
| *UpdatesWithBlankValues* | 1289 | 62336 | 109.0063ms |
| UpdatesWithBlankValues | 1124 | 58648 | 135.0077ms |
| Diffs |  :snail: 165 | :snail: 3688 | :zap: 26.0014ms |
| *UpdatesTableWithIgnoredValues* | 435 | 16728 | 147.0084ms |
| UpdatesTableWithIgnoredValues | 527 | 27680 | 148.0084ms |
| Diffs |  :zap: 92 | :zap: 10952 | :zap: 1ms |
| *UpdateDecodeVirtualAttributes* | 1046 | 54312 | 142.0081ms |
| UpdateDecodeVirtualAttributes | 929 | 51008 | 158.009ms |
| Diffs |  :snail: 117 | :snail: 3304 | :zap: 16.0009ms |
| *ToDBNameGenerateFriendlyName* | 121 | 5344 | nothing. |
| ToDBNameGenerateFriendlyName | 124 | 5440 | 1.0001ms |
| Diffs |  :zap: 3 | :zap: 96 | :zap: 1.0001ms |
| *SkipSaveAssociation* | 1323 | 55952 | 467.0267ms |
| SkipSaveAssociation | 1365 | 74680 | 475.0272ms |
| Diffs |  :zap: 42 | :zap: 18728 | :zap: 8.0005ms |
| TOTAL (original) | 610226 | 49229472 | 57.7023005s | 
| TOTAL (new) | 618783 | 35989904 | 54.1960998s |
| TOTAL (diffs) | 18446744073709543059 | 13239568 | 3.5062007s | 