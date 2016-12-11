## side by side gorm

Some tests to check gorm differences after refactoring

Example (test run on 11th of December 2016) of result produced:

| Test name | Allocs | Bytes | Duration  |
| :-------: | -----: | ----: | --------: 
| *OpenTestConnection* | 57 | 4448 | 1ms |
| OpenTestConnection | 58 | 4608 | nothing. |
| Diffs |  :zap: 1 | :zap: 160 | :snail: 1ms |
| *RunNewMigration* | 8817 | 3058064 | 4.9812849s |
| RunMigration | 8319 | 1273992 | 5.3453057s |
| Diffs |  :snail: 498 | :snail: 1784072 | :zap: 364.0208ms |
| *StringPrimaryKey* | 610 | 24200 | 259.0148ms |
| StringPrimaryKey | 643 | 36352 | 290.0166ms |
| Diffs |  :zap: 33 | :zap: 12152 | :zap: 31.0018ms |
| *SetTable* | 19456 | 983520 | 864.0494ms |
| SetTable | 19151 | 1579296 | 954.0546ms |
| Diffs |  :snail: 305 | :zap: 595776 | :zap: 90.0052ms |
| *ExceptionsWithInvalidSql* | 1390 | 78752 | 1ms |
| ExceptionsWithInvalidSql | 2331 | 1101272 | 1.0001ms |
| Diffs |  :zap: 941 | :zap: 1022520 | :zap: 100ns |
| *HasTable* | 278 | 10760 | 147.0084ms |
| HasTable | 285 | 18400 | 407.0233ms |
| Diffs |  :zap: 7 | :zap: 7640 | :zap: 260.0149ms |
| *TableName* | 186 | 12544 | 1ms |
| TableName | 161 | 22432 | nothing. |
| Diffs |  :snail: 25 | :zap: 9888 | :snail: 1ms |
| *NullValues* | 1459 | 59936 | 334.0191ms |
| NullValues | 1894 | 481712 | 350.02ms |
| Diffs |  :zap: 435 | :zap: 421776 | :zap: 16.0009ms |
| *NullValuesWithFirstOrCreate* | 1191 | 60024 | 141.0081ms |
| NullValuesWithFirstOrCreate | 967 | 55704 | 165.0094ms |
| Diffs |  :snail: 224 | :snail: 4320 | :zap: 24.0013ms |
| *Transaction* | 4210 | 217440 | 71.0041ms |
| Transaction | 4264 | 631424 | 97.0055ms |
| Diffs |  :zap: 54 | :zap: 413984 | :zap: 26.0014ms |
| *Row* | 2385 | 124496 | 216.0124ms |
| Row | 2409 | 149016 | 217.0124ms |
| Diffs |  :zap: 24 | :zap: 24520 | :zap: 1ms |
| *Rows* | 2403 | 124880 | 400.0229ms |
| Rows | 2419 | 147216 | 242.0138ms |
| Diffs |  :zap: 16 | :zap: 22336 | :snail: 158.0091ms |
| *ScanRows* | 2529 | 131000 | 177.0101ms |
| ScanRows | 2534 | 154488 | 223.0128ms |
| Diffs |  :zap: 5 | :zap: 23488 | :zap: 46.0027ms |
| *Scan* | 2747 | 142576 | 186.0106ms |
| Scan | 2932 | 183888 | 212.0121ms |
| Diffs |  :zap: 185 | :zap: 41312 | :zap: 26.0015ms |
| *Raw* | 2939 | 154728 | 218.0125ms |
| Raw | 3138 | 194368 | 290.0166ms |
| Diffs |  :zap: 199 | :zap: 39640 | :zap: 72.0041ms |
| *Group* | 170 | 5872 | nothing. |
| Group | 161 | 6240 | nothing. |
| Diffs |  :snail: 9 | :zap: 368 | :zzz: |
| *Joins* | 3927 | 228256 | 80.0045ms |
| Joins | 4195 | 275128 | 89.0051ms |
| Diffs |  :zap: 268 | :zap: 46872 | :zap: 9.0006ms |
| *JoinsWithSelect* | 1225 | 57224 | 172.0099ms |
| JoinsWithSelect | 1369 | 86096 | 119.0068ms |
| Diffs |  :zap: 144 | :zap: 28872 | :snail: 53.0031ms |
| *Having* | 118 | 5776 | nothing. |
| Having | 200 | 13352 | nothing. |
| Diffs |  :zap: 82 | :zap: 7576 | :zzz: |
| *TimeWithZone* | 4051 | 274304 | 160.0092ms |
| TimeWithZone | 3880 | 282720 | 164.0094ms |
| Diffs |  :snail: 171 | :zap: 8416 | :zap: 4.0002ms |
| *Hstore* | 27 | 1104 | nothing. |
| Hstore | 30 | 1200 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :zzz: |
| *SetAndGet* | 23 | 1184 | nothing. |
| SetAndGet | 27 | 1600 | nothing. |
| Diffs |  :zap: 4 | :zap: 416 | :zzz: |
| *CompatibilityMode* | 747 | 52632 | 1ms |
| CompatibilityMode | 527 | 35320 | nothing. |
| Diffs |  :snail: 220 | :snail: 17312 | :snail: 1ms |
| *OpenExistingDB* | 1170 | 61808 | 65.0037ms |
| OpenExistingDB | 1111 | 67672 | 75.0043ms |
| Diffs |  :snail: 59 | :zap: 5864 | :zap: 10.0006ms |
| *DdlErrors* | 266 | 15176 | 1.0001ms |
| DdlErrors | 565 | 409640 | nothing. |
| Diffs |  :zap: 299 | :zap: 394464 | :snail: 1.0001ms |
| *OpenWithOneParameter* | 20 | 864 | nothing. |
| OpenWithOneParameter | 23 | 976 | nothing. |
| Diffs |  :zap: 3 | :zap: 112 | :zzz: |
| *BelongsTo* | 10599 | 571184 | 523.03ms |
| BelongsTo | 11525 | 734072 | 682.039ms |
| Diffs |  :zap: 926 | :zap: 162888 | :zap: 159.009ms |
| *BelongsToOverrideForeignKey1* | 348 | 16928 | nothing. |
| BelongsToOverrideForeignKey1 | 341 | 20120 | 1ms |
| Diffs |  :snail: 7 | :zap: 3192 | :zap: 1ms |
| *BelongsToOverrideForeignKey2* | 278 | 13800 | nothing. |
| BelongsToOverrideForeignKey2 | 247 | 17528 | nothing. |
| Diffs |  :snail: 31 | :zap: 3728 | :zzz: |
| *HasOne* | 15537 | 842952 | 675.0386ms |
| HasOne | 15686 | 951984 | 684.0392ms |
| Diffs |  :zap: 149 | :zap: 109032 | :zap: 9.0006ms |
| *HasOneOverrideForeignKey1* | 305 | 19992 | nothing. |
| HasOneOverrideForeignKey1 | 273 | 18248 | nothing. |
| Diffs |  :snail: 32 | :snail: 1744 | :zzz: |
| *HasOneOverrideForeignKey2* | 270 | 13336 | nothing. |
| HasOneOverrideForeignKey2 | 247 | 17672 | nothing. |
| Diffs |  :snail: 23 | :zap: 4336 | :zzz: |
| *HasMany* | 11567 | 647576 | 640.0366ms |
| Many | 12088 | 811232 | 699.04ms |
| Diffs |  :zap: 521 | :zap: 163656 | :zap: 59.0034ms |
| *HasManyOverrideForeignKey1* | 300 | 14912 | nothing. |
| HasManyOverrideForeignKey1 | 269 | 17808 | nothing. |
| Diffs |  :snail: 31 | :zap: 2896 | :zzz: |
| *HasManyOverrideForeignKey2* | 267 | 14624 | 1.0001ms |
| HasManyOverrideForeignKey2 | 243 | 18688 | nothing. |
| Diffs |  :snail: 24 | :zap: 4064 | :snail: 1.0001ms |
| *ManyToMany* | 25215 | 1348832 | 1.7921025s |
| ManyToMany | 27585 | 1717240 | 1.9211098s |
| Diffs |  :zap: 2370 | :zap: 368408 | :zap: 129.0073ms |
| *Related* | 7765 | 406344 | 91.0052ms |
| Related | 7412 | 439096 | 97.0056ms |
| Diffs |  :snail: 353 | :zap: 32752 | :zap: 6.0004ms |
| *ForeignKey* | 53 | 4672 | nothing. |
| ForeignKey | 60 | 6896 | nothing. |
| Diffs |  :zap: 7 | :zap: 2224 | :zzz: |
| *LongForeignKey* | 23 | 992 | nothing. |
| LongForeignKey | 26 | 1056 | nothing. |
| Diffs |  :zap: 3 | :zap: 64 | :zzz: |
| *LongForeignKeyWithShortDest* | 23 | 1008 | nothing. |
| LongForeignKeyWithShortDest | 26 | 1072 | nothing. |
| Diffs |  :zap: 3 | :zap: 64 | :zzz: |
| *HasManyChildrenWithOneStruct* | 707 | 31232 | 72.0041ms |
| HasManyChildrenWithOneStruct | 666 | 43608 | 92.0052ms |
| Diffs |  :snail: 41 | :zap: 12376 | :zap: 20.0011ms |
| *RunCallbacks* | 2802 | 134736 | 183.0105ms |
| RunCallbacks | 2773 | 149272 | 207.0118ms |
| Diffs |  :snail: 29 | :zap: 14536 | :zap: 24.0013ms |
| *CallbacksWithErrors* | 5327 | 245112 | 176.0101ms |
| CallbacksWithErrors | 8801 | 4308408 | 234.0134ms |
| Diffs |  :zap: 3474 | :zap: 4063296 | :zap: 58.0033ms |
| *Create* | 2622 | 140536 | 107.0061ms |
| Create | 2116 | 111824 | 140.008ms |
| Diffs |  :snail: 506 | :snail: 28712 | :zap: 33.0019ms |
| *CreateWithAutoIncrement* | 31 | 1744 | nothing. |
| CreateWithAutoIncrement | 33 | 1792 | nothing. |
| Diffs |  :zap: 2 | :zap: 48 | :zzz: |
| *CreateWithNoGORMPrimayKey* | 274 | 11960 | 59.0034ms |
| CreateWithNoGORMPrimayKey | 279 | 18584 | 57.0033ms |
| Diffs |  :zap: 5 | :zap: 6624 | :snail: 2.0001ms |
| *CreateWithNoStdPrimaryKeyAndDefaultValues* | 1092 | 49880 | 127.0072ms |
| CreateWithNoStdPrimaryKeyAndDefaultValues | 1187 | 75880 | 158.0091ms |
| Diffs |  :zap: 95 | :zap: 26000 | :zap: 31.0019ms |
| *AnonymousScanner* | 1151 | 59696 | 58.0033ms |
| AnonymousScanner | 1097 | 63000 | 76.0043ms |
| Diffs |  :snail: 54 | :zap: 3304 | :zap: 18.001ms |
| *AnonymousField* | 1653 | 84680 | 69.004ms |
| AnonymousField | 1620 | 96472 | 89.0051ms |
| Diffs |  :snail: 33 | :zap: 11792 | :zap: 20.0011ms |
| *SelectWithCreate* | 3103 | 150808 | 145.0083ms |
| SelectWithCreate | 3253 | 205624 | 164.0094ms |
| Diffs |  :zap: 150 | :zap: 54816 | :zap: 19.0011ms |
| *OmitWithCreate* | 3285 | 167696 | 120.0068ms |
| OmitWithCreate | 3423 | 217224 | 163.0093ms |
| Diffs |  :zap: 138 | :zap: 49528 | :zap: 43.0025ms |
| *CustomizeColumn* | 903 | 42032 | 307.0176ms |
| CustomizeColumn | 861 | 59328 | 323.0184ms |
| Diffs |  :snail: 42 | :zap: 17296 | :zap: 16.0008ms |
| *CustomColumnAndIgnoredFieldClash* | 161 | 13896 | 150.0086ms |
| CustomColumnAndIgnoredFieldClash | 160 | 10488 | 150.0086ms |
| Diffs |  :snail: 1 | :snail: 3408 | :zzz: |
| *ManyToManyWithCustomizedColumn* | 1688 | 77184 | 651.0372ms |
| ManyToManyWithCustomizedColumn | 2082 | 138920 | 656.0375ms |
| Diffs |  :zap: 394 | :zap: 61736 | :zap: 5.0003ms |
| *OneToOneWithCustomizedColumn* | 1575 | 74960 | 584.0334ms |
| OneToOneWithCustomizedColumn | 1563 | 98152 | 675.0386ms |
| Diffs |  :snail: 12 | :zap: 23192 | :zap: 91.0052ms |
| *OneToManyWithCustomizedColumn* | 3367 | 167616 | 517.0296ms |
| OneToManyWithCustomizedColumn | 3510 | 217408 | 625.0357ms |
| Diffs |  :zap: 143 | :zap: 49792 | :zap: 108.0061ms |
| *HasOneWithPartialCustomizedColumn* | 2318 | 113488 | 582.0333ms |
| HasOneWithPartialCustomizedColumn | 2447 | 148472 | 651.0372ms |
| Diffs |  :zap: 129 | :zap: 34984 | :zap: 69.0039ms |
| *BelongsToWithPartialCustomizedColumn* | 2550 | 127208 | 569.0325ms |
| BelongsToWithPartialCustomizedColumn | 2691 | 166624 | 689.0394ms |
| Diffs |  :zap: 141 | :zap: 39416 | :zap: 120.0069ms |
| *Delete* | 2298 | 120752 | 216.0124ms |
| Delete | 2181 | 126608 | 223.0128ms |
| Diffs |  :snail: 117 | :zap: 5856 | :zap: 7.0004ms |
| *InlineDelete* | 2316 | 121584 | 260.0148ms |
| InlineDelete | 2303 | 136744 | 298.0171ms |
| Diffs |  :snail: 13 | :zap: 15160 | :zap: 38.0023ms |
| *SoftDelete* | 1042 | 42584 | 234.0134ms |
| SoftDelete | 1282 | 78184 | 274.0157ms |
| Diffs |  :zap: 240 | :zap: 35600 | :zap: 40.0023ms |
| *PrefixColumnNameForEmbeddedStruct* | 435 | 19920 | 1ms |
| PrefixColumnNameForEmbeddedStruct | 424 | 30952 | nothing. |
| Diffs |  :snail: 11 | :zap: 11032 | :snail: 1ms |
| *SaveAndQueryEmbeddedStruct* | 1288 | 52304 | 185.0106ms |
| SaveAndQueryEmbeddedStruct | 1368 | 71312 | 282.0162ms |
| Diffs |  :zap: 80 | :zap: 19008 | :zap: 97.0056ms |
| *CalculateField* | 493 | 24968 | 1ms |
| CalculateField | 450 | 31544 | nothing. |
| Diffs |  :snail: 43 | :zap: 6576 | :snail: 1ms |
| *JoinTable* | 4055 | 202176 | 507.029ms |
| JoinTable | 4326 | 284128 | 515.0295ms |
| Diffs |  :zap: 271 | :zap: 81952 | :zap: 8.0005ms |
| *Indexes* | 8220 | 392256 | 713.0408ms |
| Indexes | 9815 | 3085472 | 1.1130636s |
| Diffs |  :zap: 1595 | :zap: 2693216 | :zap: 400.0228ms |
| *AutoMigration* | 1347 | 52824 | 330.0189ms |
| AutoMigration | 1345 | 59680 | 462.0264ms |
| Diffs |  :snail: 2 | :zap: 6856 | :zap: 132.0075ms |
| *MultipleIndexes* | 2438 | 101744 | 812.0464ms |
| MultipleIndexes | 3086 | 939360 | 878.0502ms |
| Diffs |  :zap: 648 | :zap: 837616 | :zap: 66.0038ms |
| *ManyToManyWithMultiPrimaryKeys* | 22 | 1040 | nothing. |
| ManyToManyWithMultiPrimaryKeys | 25 | 1136 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :zzz: |
| *ManyToManyWithCustomizedForeignKeys* | 22 | 1056 | nothing. |
| ManyToManyWithCustomizedForeignKeys | 25 | 1152 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :zzz: |
| *ManyToManyWithCustomizedForeignKeys2* | 22 | 1056 | 1.0001ms |
| ManyToManyWithCustomizedForeignKeys2 | 25 | 1152 | nothing. |
| Diffs |  :zap: 3 | :zap: 96 | :snail: 1.0001ms |
| *PointerFields* | 2046 | 82856 | 398.0227ms |
| PointerFields | 2562 | 156488 | 508.029ms |
| Diffs |  :zap: 516 | :zap: 73632 | :zap: 110.0063ms |
| *Polymorphic* | 17034 | 868336 | 1.1570662s |
| Polymorphic | 23325 | 1611192 | 1.5580891s |
| Diffs |  :zap: 6291 | :zap: 742856 | :zap: 401.0229ms |
| *NamedPolymorphic* | 11499 | 626368 | 978.056ms |
| NamedPolymorphic | 16185 | 1130040 | 1.1220642s |
| Diffs |  :zap: 4686 | :zap: 503672 | :zap: 144.0082ms |
| *Preload* | 22326 | 1072608 | 386.022ms |
| Preload | 22744 | 1309208 | 437.025ms |
| Diffs |  :zap: 418 | :zap: 236600 | :zap: 51.003ms |
| *NestedPreload1* | 1760 | 116008 | 495.0284ms |
| NestedPreload1 | 1958 | 124472 | 642.0367ms |
| Diffs |  :zap: 198 | :zap: 8464 | :zap: 147.0083ms |
| *NestedPreload2* | 2215 | 93432 | 574.0328ms |
| NestedPreload2 | 2389 | 144536 | 759.0434ms |
| Diffs |  :zap: 174 | :zap: 51104 | :zap: 185.0106ms |
| *NestedPreload3* | 1978 | 90344 | 458.0262ms |
| NestedPreload3 | 2146 | 132192 | 634.0363ms |
| Diffs |  :zap: 168 | :zap: 41848 | :zap: 176.0101ms |
| *NestedPreload4* | 1758 | 74464 | 549.0314ms |
| NestedPreload4 | 1917 | 120624 | 601.0344ms |
| Diffs |  :zap: 159 | :zap: 46160 | :zap: 52.003ms |
| *NestedPreload5* | 2194 | 90312 | 629.0359ms |
| NestedPreload5 | 2376 | 143496 | 903.0517ms |
| Diffs |  :zap: 182 | :zap: 53184 | :zap: 274.0158ms |
| *NestedPreload6* | 3475 | 140488 | 633.0362ms |
| NestedPreload6 | 3688 | 220792 | 783.0447ms |
| Diffs |  :zap: 213 | :zap: 80304 | :zap: 150.0085ms |
| *NestedPreload7* | 3077 | 127976 | 744.0426ms |
| NestedPreload7 | 3290 | 191320 | 697.0398ms |
| Diffs |  :zap: 213 | :zap: 63344 | :snail: 47.0028ms |
| *NestedPreload8* | 2614 | 106328 | 566.0324ms |
| NestedPreload8 | 2807 | 166872 | 734.042ms |
| Diffs |  :zap: 193 | :zap: 60544 | :zap: 168.0096ms |
| *NestedPreload9* | 5991 | 251696 | 841.0481ms |
| NestedPreload9 | 6275 | 383168 | 1.084062s |
| Diffs |  :zap: 284 | :zap: 131472 | :zap: 243.0139ms |
| *NestedPreload10* | 2111 | 105112 | 671.0384ms |
| NestedPreload10 | 2261 | 133008 | 869.0497ms |
| Diffs |  :zap: 150 | :zap: 27896 | :zap: 198.0113ms |
| *NestedPreload11* | 1808 | 77528 | 686.0393ms |
| NestedPreload11 | 2020 | 123064 | 731.0418ms |
| Diffs |  :zap: 212 | :zap: 45536 | :zap: 45.0025ms |
| *NestedPreload12* | 2491 | 117360 | 655.0375ms |
| NestedPreload12 | 2689 | 158880 | 834.0477ms |
| Diffs |  :zap: 198 | :zap: 41520 | :zap: 179.0102ms |
| *ManyToManyPreloadWithMultiPrimaryKeys* | 24 | 14832 | nothing. |
| ManyToManyPreloadWithMultiPrimaryKeys | 25 | 1152 | nothing. |
| Diffs |  :zap: 1 | :snail: 13680 | :zzz: |
| *ManyToManyPreloadForNestedPointer* | 6491 | 285848 | 716.041ms |
| ManyToManyPreloadForNestedPointer | 8603 | 544704 | 970.0555ms |
| Diffs |  :zap: 2112 | :zap: 258856 | :zap: 254.0145ms |
| *NestedManyToManyPreload* | 4159 | 182672 | 864.0494ms |
| NestedManyToManyPreload | 5374 | 365400 | 1.0170582s |
| Diffs |  :zap: 1215 | :zap: 182728 | :zap: 153.0088ms |
| *NestedManyToManyPreload2* | 2646 | 120752 | 616.0353ms |
| NestedManyToManyPreload2 | 3319 | 215928 | 775.0443ms |
| Diffs |  :zap: 673 | :zap: 95176 | :zap: 159.009ms |
| *NestedManyToManyPreload3* | 4437 | 190224 | 741.0424ms |
| NestedManyToManyPreload3 | 5402 | 350392 | 959.0548ms |
| Diffs |  :zap: 965 | :zap: 160168 | :zap: 218.0124ms |
| *NestedManyToManyPreload3ForStruct* | 4653 | 199136 | 888.0508ms |
| NestedManyToManyPreload3ForStruct | 5632 | 360880 | 1.0770616s |
| Diffs |  :zap: 979 | :zap: 161744 | :zap: 189.0108ms |
| *NestedManyToManyPreload4* | 3471 | 151360 | 1.1220642s |
| NestedManyToManyPreload4 | 4212 | 294056 | 1.0950626s |
| Diffs |  :zap: 741 | :zap: 142696 | :snail: 27.0016ms |
| *ManyToManyPreloadForPointer* | 4836 | 224896 | 606.0347ms |
| ManyToManyPreloadForPointer | 6613 | 430072 | 720.0411ms |
| Diffs |  :zap: 1777 | :zap: 205176 | :zap: 114.0064ms |
| *NilPointerSlice* | 1847 | 76504 | 556.0318ms |
| NilPointerSlice | 1996 | 120296 | 733.0419ms |
| Diffs |  :zap: 149 | :zap: 43792 | :zap: 177.0101ms |
| *NilPointerSlice2* | 1714 | 74848 | 822.047ms |
| NilPointerSlice2 | 1847 | 125264 | 985.0564ms |
| Diffs |  :zap: 133 | :zap: 50416 | :zap: 163.0094ms |
| *PrefixedPreloadDuplication* | 4024 | 164792 | 1.0650609s |
| PrefixedPreloadDuplication | 4306 | 253616 | 1.2530717s |
| Diffs |  :zap: 282 | :zap: 88824 | :zap: 188.0108ms |
| *FirstAndLast* | 4609 | 244120 | 178.0102ms |
| FirstAndLast | 3911 | 217288 | 223.0127ms |
| Diffs |  :snail: 698 | :snail: 26832 | :zap: 45.0025ms |
| *FirstAndLastWithNoStdPrimaryKey* | 1549 | 72208 | 132.0076ms |
| FirstAndLastWithNoStdPrimaryKey | 1582 | 96160 | 150.0085ms |
| Diffs |  :zap: 33 | :zap: 23952 | :zap: 18.0009ms |
| *UIntPrimaryKey* | 566 | 28352 | 1.0001ms |
| UIntPrimaryKey | 485 | 28600 | nothing. |
| Diffs |  :snail: 81 | :zap: 248 | :snail: 1.0001ms |
| *StringPrimaryKeyForNumericValueStartingWithZero* | 493 | 20512 | 1.0001ms |
| StringPrimaryKeyForNumericValueStartingWithZero | 921 | 431088 | 1ms |
| Diffs |  :zap: 428 | :zap: 410576 | :snail: 100ns |
| *FindAsSliceOfPointers* | 20581 | 1289328 | 85.0049ms |
| FindAsSliceOfPointers | 15639 | 895632 | 98.0056ms |
| Diffs |  :snail: 4942 | :snail: 393696 | :zap: 13.0007ms |
| *SearchWithPlainSQL* | 10386 | 656952 | 238.0136ms |
| SearchWithPlainSQL | 10104 | 654008 | 270.0154ms |
| Diffs |  :snail: 282 | :snail: 2944 | :zap: 32.0018ms |
| *SearchWithStruct* | 7624 | 434928 | 232.0133ms |
| SearchWithStruct | 6312 | 353128 | 258.0148ms |
| Diffs |  :snail: 1312 | :snail: 81800 | :zap: 26.0015ms |
| *SearchWithMap* | 6148 | 337976 | 313.0179ms |
| SearchWithMap | 5297 | 309728 | 354.0202ms |
| Diffs |  :snail: 851 | :snail: 28248 | :zap: 41.0023ms |
| *SearchWithEmptyChain* | 4177 | 225920 | 224.0128ms |
| SearchWithEmptyChain | 3984 | 234480 | 225.0129ms |
| Diffs |  :snail: 193 | :zap: 8560 | :zap: 1.0001ms |
| *Select* | 1054 | 55656 | 71.0041ms |
| Select | 1006 | 58160 | 70.004ms |
| Diffs |  :snail: 48 | :zap: 2504 | :snail: 1.0001ms |
| *OrderAndPluck* | 15447 | 957744 | 229.0131ms |
| OrderAndPluck | 12060 | 702216 | 240.0137ms |
| Diffs |  :snail: 3387 | :snail: 255528 | :zap: 11.0006ms |
| *Limit* | 20281 | 1357360 | 356.0204ms |
| Limit | 15886 | 1042904 | 386.022ms |
| Diffs |  :snail: 4395 | :snail: 314456 | :zap: 30.0016ms |
| *Offset* | 88228 | 5791144 | 1.3240757s |
| Offset | 68949 | 4363736 | 1.626093s |
| Diffs |  :snail: 19279 | :snail: 1427408 | :zap: 302.0173ms |
| *Or* | 2509 | 153152 | 212.0121ms |
| Or | 2436 | 148552 | 242.0139ms |
| Diffs |  :snail: 73 | :snail: 4600 | :zap: 30.0018ms |
| *Count* | 3248 | 176728 | 225.0128ms |
| Count | 3423 | 208792 | 242.0139ms |
| Diffs |  :zap: 175 | :zap: 32064 | :zap: 17.0011ms |
| *Not* | 22104 | 1189368 | 598.0342ms |
| Not | 21565 | 1529808 | 571.0326ms |
| Diffs |  :snail: 539 | :zap: 340440 | :snail: 27.0016ms |
| *FillSmallerStruct* | 912 | 42584 | 55.0032ms |
| FillSmallerStruct | 958 | 56112 | 75.0043ms |
| Diffs |  :zap: 46 | :zap: 13528 | :zap: 20.0011ms |
| *FindOrInitialize* | 7206 | 408664 | 55.0031ms |
| FindOrInitialize | 5244 | 276768 | 111.0064ms |
| Diffs |  :snail: 1962 | :snail: 131896 | :zap: 56.0033ms |
| *FindOrCreate* | 11997 | 646736 | 506.0289ms |
| FindOrCreate | 10423 | 1332888 | 494.0282ms |
| Diffs |  :snail: 1574 | :zap: 686152 | :snail: 12.0007ms |
| *SelectWithEscapedFieldName* | 2263 | 117296 | 177.0102ms |
| SelectWithEscapedFieldName | 2053 | 122264 | 230.0132ms |
| Diffs |  :snail: 210 | :zap: 4968 | :zap: 53.003ms |
| *SelectWithVariables* | 684 | 34160 | 61.0035ms |
| SelectWithVariables | 655 | 39416 | 80.0045ms |
| Diffs |  :snail: 29 | :zap: 5256 | :zap: 19.001ms |
| *FirstAndLastWithRaw* | 2663 | 135904 | 150.0086ms |
| FirstAndLastWithRaw | 2548 | 147936 | 208.0119ms |
| Diffs |  :snail: 115 | :zap: 12032 | :zap: 58.0033ms |
| *ScannableSlices* | 2948 | 144560 | 80.0045ms |
| ScannableSlices | 996 | 51656 | 78.0045ms |
| Diffs |  :snail: 1952 | :snail: 92904 | :snail: 2ms |
| *Scopes* | 3642 | 204256 | 184.0105ms |
| Scopes | 3462 | 209200 | 231.0132ms |
| Diffs |  :snail: 180 | :zap: 4944 | :zap: 47.0027ms |
| *Update* | 6851 | 322880 | 492.0281ms |
| Update | 6502 | 346752 | 574.0329ms |
| Diffs |  :snail: 349 | :zap: 23872 | :zap: 82.0048ms |
| *UpdateWithNoStdPrimaryKeyAndDefaultValues* | 2992 | 134448 | 518.0296ms |
| UpdateWithNoStdPrimaryKeyAndDefaultValues | 3012 | 169560 | 607.0347ms |
| Diffs |  :zap: 20 | :zap: 35112 | :zap: 89.0051ms |
| *Updates* | 4788 | 217952 | 318.0182ms |
| Updates | 4599 | 241288 | 332.019ms |
| Diffs |  :snail: 189 | :zap: 23336 | :zap: 14.0008ms |
| *UpdateColumn* | 3265 | 149912 | 305.0174ms |
| UpdateColumn | 2849 | 145768 | 345.0198ms |
| Diffs |  :snail: 416 | :snail: 4144 | :zap: 40.0024ms |
| *SelectWithUpdate* | 7174 | 341664 | 222.0127ms |
| SelectWithUpdate | 7263 | 443160 | 305.0174ms |
| Diffs |  :zap: 89 | :zap: 101496 | :zap: 83.0047ms |
| *SelectWithUpdateWithMap* | 7210 | 345048 | 247.0142ms |
| SelectWithUpdateWithMap | 7287 | 442960 | 286.0163ms |
| Diffs |  :zap: 77 | :zap: 97912 | :zap: 39.0021ms |
| *OmitWithUpdate* | 6039 | 292464 | 246.0141ms |
| OmitWithUpdate | 6123 | 371912 | 262.015ms |
| Diffs |  :zap: 84 | :zap: 79448 | :zap: 16.0009ms |
| *OmitWithUpdateWithMap* | 5878 | 287568 | 182.0104ms |
| OmitWithUpdateWithMap | 5979 | 366464 | 209.012ms |
| Diffs |  :zap: 101 | :zap: 78896 | :zap: 27.0016ms |
| *SelectWithUpdateColumn* | 4392 | 213008 | 174.0099ms |
| SelectWithUpdateColumn | 4072 | 240536 | 192.011ms |
| Diffs |  :snail: 320 | :zap: 27528 | :zap: 18.0011ms |
| *OmitWithUpdateColumn* | 4393 | 213584 | 158.0091ms |
| OmitWithUpdateColumn | 4067 | 240088 | 175.01ms |
| Diffs |  :snail: 326 | :zap: 26504 | :zap: 17.0009ms |
| *UpdateColumnsSkipsAssociations* | 4328 | 208096 | 258.0147ms |
| UpdateColumnsSkipsAssociations | 4044 | 237448 | 249.0142ms |
| Diffs |  :snail: 284 | :zap: 29352 | :snail: 9.0005ms |
| *UpdatesWithBlankValues* | 1289 | 61840 | 142.0082ms |
| UpdatesWithBlankValues | 1124 | 58648 | 174.0099ms |
| Diffs |  :snail: 165 | :snail: 3192 | :zap: 32.0017ms |
| *UpdatesTableWithIgnoredValues* | 435 | 16728 | 126.0072ms |
| UpdatesTableWithIgnoredValues | 527 | 27680 | 157.009ms |
| Diffs |  :zap: 92 | :zap: 10952 | :zap: 31.0018ms |
| *UpdateDecodeVirtualAttributes* | 1045 | 54264 | 126.0072ms |
| UpdateDecodeVirtualAttributes | 929 | 51008 | 141.0081ms |
| Diffs |  :snail: 116 | :snail: 3256 | :zap: 15.0009ms |
| *ToDBNameGenerateFriendlyName* | 120 | 5056 | nothing. |
| ToDBNameGenerateFriendlyName | 124 | 5424 | nothing. |
| Diffs |  :zap: 4 | :zap: 368 | :zzz: |
| *SkipSaveAssociation* | 1323 | 55952 | 400.0229ms |
| SkipSaveAssociation | 1365 | 74680 | 442.0253ms |
| Diffs |  :zap: 42 | :zap: 18728 | :zap: 42.0024ms |
| TOTAL (original) | 610564 | 49250336 | 55.9141975s |
| TOTAL (new) | 619078 | 36010600 | 47.953743s |
| TOTAL (diffs) |  :snail: 8514 |   :zap: 13239736 |   :zap: 7.9604545s |