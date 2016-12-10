package tests

import (
	"SideBySideGorm/new_types"
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"testing"
	"time"
)

type (
	MeasureData []*Measure
	Measure     struct {
		pairId      int
		name        string
		duration    uint64
		startAllocs uint64 // The initial states of memStats.Mallocs and memStats.TotalAlloc.
		startBytes  uint64
		netAllocs   uint64 // The net total of this test after being run.
		netBytes    uint64
		start       time.Time
		pair        *Measure
		isNew       bool
	}
)

var (
	memStats     runtime.MemStats
	measuresData MeasureData = make(MeasureData, 0, 0)
)

func measureAndRun(t *testing.T, name string, f func(t *testing.T)) bool {

	nameParts := strings.Split(name, " ")
	pairId, err := strconv.Atoi(nameParts[0])
	if err != nil {
		t.Fatalf("ERROR : %v", err)
	}
	measurement := &Measure{
		pairId: pairId,
		name:   nameParts[1],
		isNew:  true,
	}

	for _, pairMeas := range measuresData {
		if pairMeas.pairId == pairId {
			measurement.isNew = false
			pairMeas.pair = measurement
			break
		}
	}
	//t.Logf("Processing %s with id %d (is new ? %t)", nameParts[1], pairId, measurement.isNew)

	runtime.ReadMemStats(&memStats)
	measurement.startAllocs = memStats.Mallocs
	measurement.startBytes = memStats.TotalAlloc

	measurement.start = time.Now()
	result := t.Run(name, f)
	measurement.duration = uint64(time.Now().Sub(measurement.start).Nanoseconds())

	runtime.ReadMemStats(&memStats)

	measurement.netAllocs += memStats.Mallocs - measurement.startAllocs
	measurement.netBytes += memStats.TotalAlloc - measurement.startBytes

	measuresData = append(measuresData, measurement)

	return result
}

func TestEverything(t *testing.T) {
	measureAndRun(t, "0 OpenTestConnection", OpenNewTestConnection)
	if new_types.TestDBErr != nil {
		t.Fatalf("No error should happen when connecting to test database, but got err=%+v", new_types.TestDBErr)
	}
	measureAndRun(t, "0 OpenTestConnection", OpenTestConnection)
	if oldDBError != nil {
		t.Fatalf("No error should happen when connecting to old test database, but got err=%+v", oldDBError)
	}

	measureAndRun(t, "1 RunNewMigration", RunNewMigration)
	measureAndRun(t, "1 RunMigration", OldRunMigration)

	measureAndRun(t, "2 StringPrimaryKey", StringPrimaryKey)
	measureAndRun(t, "2 StringPrimaryKey", OldStringPrimaryKey)

	measureAndRun(t, "3 SetTable", SetTable)
	measureAndRun(t, "3 SetTable", OldSetTable)

	measureAndRun(t, "4 ExceptionsWithInvalidSql", ExceptionsWithInvalidSql)
	measureAndRun(t, "4 ExceptionsWithInvalidSql", OldExceptionsWithInvalidSql)

	measureAndRun(t, "5 HasTable", HasTable)
	measureAndRun(t, "5 HasTable", OldHasTable)

	measureAndRun(t, "6 TableName", TableName)
	measureAndRun(t, "6 TableName", OldTableName)

	measureAndRun(t, "7 NullValues", NullValues)
	measureAndRun(t, "7 NullValues", OldNullValues)

	measureAndRun(t, "8 NullValuesWithFirstOrCreate", NullValuesWithFirstOrCreate)
	measureAndRun(t, "8 NullValuesWithFirstOrCreate", OldNullValuesWithFirstOrCreate)

	measureAndRun(t, "9 Transaction", Transaction)
	measureAndRun(t, "9 Transaction", OldTransaction)

	measureAndRun(t, "10 Row", Row)
	measureAndRun(t, "10 Row", OldRow)

	measureAndRun(t, "11 Rows", Rows)
	measureAndRun(t, "11 Rows", OldRows)

	measureAndRun(t, "12 ScanRows", ScanRows)
	measureAndRun(t, "12 ScanRows", OldScanRows)

	measureAndRun(t, "13 Scan", Scan)
	measureAndRun(t, "13 Scan", OldScan)

	measureAndRun(t, "14 Raw", Raw)
	measureAndRun(t, "14 Raw", OldRaw)

	measureAndRun(t, "15 Group", Group)
	measureAndRun(t, "15 Group", OldGroup)

	measureAndRun(t, "16 Joins", Joins)
	measureAndRun(t, "16 Joins", OldJoins)

	measureAndRun(t, "17 JoinsWithSelect", JoinsWithSelect)
	measureAndRun(t, "17 JoinsWithSelect", OldJoinsWithSelect)

	measureAndRun(t, "18 Having", Having)
	measureAndRun(t, "18 Having", OldHaving)

	measureAndRun(t, "19 TimeWithZone", TimeWithZone)
	measureAndRun(t, "19 TimeWithZone", OldTimeWithZone)

	measureAndRun(t, "20 Hstore", Hstore)
	measureAndRun(t, "20 Hstore", OldHstore)

	measureAndRun(t, "21 SetAndGet", SetAndGet)
	measureAndRun(t, "21 SetAndGet", OldSetAndGet)

	measureAndRun(t, "22 CompatibilityMode", CompatibilityMode)
	measureAndRun(t, "22 CompatibilityMode", OldCompatibilityMode)

	measureAndRun(t, "23 OpenExistingDB", OpenExistingDB)
	measureAndRun(t, "23 OpenExistingDB", OldOpenExistingDB)

	measureAndRun(t, "24 DdlErrors", DdlErrors)
	measureAndRun(t, "24 DdlErrors", OldDdlErrors)

	measureAndRun(t, "25 OpenWithOneParameter", OpenWithOneParameter)
	measureAndRun(t, "25 OpenWithOneParameter", OldOpenWithOneParameter)

	measureAndRun(t, "26 BelongsTo", BelongsTo)
	measureAndRun(t, "26 BelongsTo", OldBelongsTo)

	measureAndRun(t, "27 BelongsToOverrideForeignKey1", BelongsToOverrideForeignKey1)
	measureAndRun(t, "27 BelongsToOverrideForeignKey1", OldBelongsToOverrideForeignKey1)

	measureAndRun(t, "28 BelongsToOverrideForeignKey2", BelongsToOverrideForeignKey2)
	measureAndRun(t, "28 BelongsToOverrideForeignKey2", OldBelongsToOverrideForeignKey2)

	measureAndRun(t, "29 HasOne", HasOne)
	measureAndRun(t, "29 HasOne", OldHasOne)

	measureAndRun(t, "30 HasOneOverrideForeignKey1", HasOneOverrideForeignKey1)
	measureAndRun(t, "30 HasOneOverrideForeignKey1", OldHasOneOverrideForeignKey1)

	measureAndRun(t, "31 HasOneOverrideForeignKey2", HasOneOverrideForeignKey2)
	measureAndRun(t, "31 HasOneOverrideForeignKey2", OldHasOneOverrideForeignKey2)

	measureAndRun(t, "32 HasMany", HasMany)
	measureAndRun(t, "32 Many", OldHasMany)

	measureAndRun(t, "33 HasManyOverrideForeignKey1", HasManyOverrideForeignKey1)
	measureAndRun(t, "33 HasManyOverrideForeignKey1", OldHasManyOverrideForeignKey1)

	measureAndRun(t, "34 HasManyOverrideForeignKey2", HasManyOverrideForeignKey2)
	measureAndRun(t, "34 HasManyOverrideForeignKey2", OldHasManyOverrideForeignKey2)

	measureAndRun(t, "35 ManyToMany", ManyToMany)
	measureAndRun(t, "35 ManyToMany", OldManyToMany)

	measureAndRun(t, "36 Related", Related)
	measureAndRun(t, "36 Related", OldRelated)

	measureAndRun(t, "37 ForeignKey", ForeignKey)
	measureAndRun(t, "37 ForeignKey", OldForeignKey)

	measureAndRun(t, "38 LongForeignKey", LongForeignKey)
	measureAndRun(t, "38 LongForeignKey", OldLongForeignKey)

	measureAndRun(t, "39 LongForeignKeyWithShortDest", LongForeignKeyWithShortDest)
	measureAndRun(t, "39 LongForeignKeyWithShortDest", OldLongForeignKeyWithShortDest)

	measureAndRun(t, "40 HasManyChildrenWithOneStruct", HasManyChildrenWithOneStruct)
	measureAndRun(t, "40 HasManyChildrenWithOneStruct", OldHasManyChildrenWithOneStruct)

	measureAndRun(t, "41 RunCallbacks", RunCallbacks)
	measureAndRun(t, "41 RunCallbacks", OldRunCallbacks)

	measureAndRun(t, "42 CallbacksWithErrors", CallbacksWithErrors)
	measureAndRun(t, "42 CallbacksWithErrors", OldCallbacksWithErrors)

	measureAndRun(t, "43 Create", Create)
	measureAndRun(t, "43 Create", OldCreate)

	measureAndRun(t, "44 CreateWithAutoIncrement", CreateWithAutoIncrement)
	measureAndRun(t, "44 CreateWithAutoIncrement", OldCreateWithAutoIncrement)

	measureAndRun(t, "45 CreateWithNoGORMPrimayKey", CreateWithNoGORMPrimayKey)
	measureAndRun(t, "45 CreateWithNoGORMPrimayKey", OldCreateWithNoGORMPrimayKey)

	measureAndRun(t, "46 CreateWithNoStdPrimaryKeyAndDefaultValues", CreateWithNoStdPrimaryKeyAndDefaultValues)
	measureAndRun(t, "46 CreateWithNoStdPrimaryKeyAndDefaultValues", OldCreateWithNoStdPrimaryKeyAndDefaultValues)

	measureAndRun(t, "47 AnonymousScanner", AnonymousScanner)
	measureAndRun(t, "47 AnonymousScanner", OldAnonymousScanner)

	measureAndRun(t, "48 AnonymousField", AnonymousField)
	measureAndRun(t, "48 AnonymousField", OldAnonymousField)

	measureAndRun(t, "49 SelectWithCreate", SelectWithCreate)
	measureAndRun(t, "49 SelectWithCreate", OldSelectWithCreate)

	measureAndRun(t, "50 OmitWithCreate", OmitWithCreate)
	measureAndRun(t, "50 OmitWithCreate", OldOmitWithCreate)

	measureAndRun(t, "51 CustomizeColumn", DoCustomizeColumn)
	measureAndRun(t, "51 CustomizeColumn", OldDoCustomizeColumn)

	measureAndRun(t, "52 CustomColumnAndIgnoredFieldClash", DoCustomColumnAndIgnoredFieldClash)
	measureAndRun(t, "52 CustomColumnAndIgnoredFieldClash", OldDoCustomColumnAndIgnoredFieldClash)

	measureAndRun(t, "53 ManyToManyWithCustomizedColumn", ManyToManyWithCustomizedColumn)
	measureAndRun(t, "53 ManyToManyWithCustomizedColumn", OldManyToManyWithCustomizedColumn)

	measureAndRun(t, "54 OneToOneWithCustomizedColumn", OneToOneWithCustomizedColumn)
	measureAndRun(t, "54 OneToOneWithCustomizedColumn", OldOneToOneWithCustomizedColumn)

	measureAndRun(t, "55 OneToManyWithCustomizedColumn", OneToManyWithCustomizedColumn)
	measureAndRun(t, "55 OneToManyWithCustomizedColumn", OldOneToManyWithCustomizedColumn)

	measureAndRun(t, "56 HasOneWithPartialCustomizedColumn", HasOneWithPartialCustomizedColumn)
	measureAndRun(t, "56 HasOneWithPartialCustomizedColumn", OldHasOneWithPartialCustomizedColumn)

	measureAndRun(t, "57 BelongsToWithPartialCustomizedColumn", BelongsToWithPartialCustomizedColumn)
	measureAndRun(t, "57 BelongsToWithPartialCustomizedColumn", OldBelongsToWithPartialCustomizedColumn)

	measureAndRun(t, "58 Delete", DoDelete)
	measureAndRun(t, "58 Delete", OldDoDelete)

	measureAndRun(t, "59 InlineDelete", InlineDelete)
	measureAndRun(t, "59 InlineDelete", OldInlineDelete)

	measureAndRun(t, "60 SoftDelete", SoftDelete)
	measureAndRun(t, "60 SoftDelete", OldSoftDelete)

	measureAndRun(t, "61 PrefixColumnNameForEmbeddedStruct", PrefixColumnNameForEmbeddedStruct)
	measureAndRun(t, "61 PrefixColumnNameForEmbeddedStruct", OldPrefixColumnNameForEmbeddedStruct)

	measureAndRun(t, "62 SaveAndQueryEmbeddedStruct", SaveAndQueryEmbeddedStruct)
	measureAndRun(t, "62 SaveAndQueryEmbeddedStruct", OldSaveAndQueryEmbeddedStruct)

	measureAndRun(t, "63 CalculateField", DoCalculateField)
	measureAndRun(t, "63 CalculateField", OldDoCalculateField)

	measureAndRun(t, "64 JoinTable", DoJoinTable)
	measureAndRun(t, "64 JoinTable", OldDoJoinTable)

	measureAndRun(t, "65 Indexes", Indexes)
	measureAndRun(t, "65 Indexes", OldIndexes)

	measureAndRun(t, "66 AutoMigration", AutoMigration)
	measureAndRun(t, "66 AutoMigration", OldAutoMigration)

	measureAndRun(t, "67 MultipleIndexes", DoMultipleIndexes)
	measureAndRun(t, "67 MultipleIndexes", OldDoMultipleIndexes)

	measureAndRun(t, "68 ManyToManyWithMultiPrimaryKeys", ManyToManyWithMultiPrimaryKeys)
	measureAndRun(t, "68 ManyToManyWithMultiPrimaryKeys", OldManyToManyWithMultiPrimaryKeys)

	measureAndRun(t, "69 ManyToManyWithCustomizedForeignKeys", ManyToManyWithCustomizedForeignKeys)
	measureAndRun(t, "69 ManyToManyWithCustomizedForeignKeys", OldManyToManyWithCustomizedForeignKeys)

	measureAndRun(t, "70 ManyToManyWithCustomizedForeignKeys2", ManyToManyWithCustomizedForeignKeys2)
	measureAndRun(t, "70 ManyToManyWithCustomizedForeignKeys2", OldManyToManyWithCustomizedForeignKeys2)

	measureAndRun(t, "71 PointerFields", PointerFields)
	measureAndRun(t, "71 PointerFields", OldPointerFields)

	measureAndRun(t, "72 Polymorphic", Polymorphic)
	measureAndRun(t, "72 Polymorphic", OldPolymorphic)

	measureAndRun(t, "73 NamedPolymorphic", NamedPolymorphic)
	measureAndRun(t, "73 NamedPolymorphic", OldNamedPolymorphic)

	measureAndRun(t, "74 Preload", Preload)
	measureAndRun(t, "74 Preload", OldPreload)

	measureAndRun(t, "75 NestedPreload1", NestedPreload1)
	measureAndRun(t, "75 NestedPreload1", OldNestedPreload1)

	measureAndRun(t, "76 NestedPreload2", NestedPreload2)
	measureAndRun(t, "76 NestedPreload2", OldNestedPreload2)

	measureAndRun(t, "77 NestedPreload3", NestedPreload3)
	measureAndRun(t, "77 NestedPreload3", OldNestedPreload3)

	measureAndRun(t, "78 NestedPreload4", NestedPreload4)
	measureAndRun(t, "78 NestedPreload4", OldNestedPreload4)

	measureAndRun(t, "79 NestedPreload5", NestedPreload5)
	measureAndRun(t, "79 NestedPreload5", OldNestedPreload5)

	measureAndRun(t, "80 NestedPreload6", NestedPreload6)
	measureAndRun(t, "80 NestedPreload6", OldNestedPreload6)

	measureAndRun(t, "81 NestedPreload7", NestedPreload7)
	measureAndRun(t, "81 NestedPreload7", OldNestedPreload7)

	measureAndRun(t, "82 NestedPreload8", NestedPreload8)
	measureAndRun(t, "82 NestedPreload8", OldNestedPreload8)

	measureAndRun(t, "83 NestedPreload9", NestedPreload9)
	measureAndRun(t, "83 NestedPreload9", OldNestedPreload9)

	measureAndRun(t, "84 NestedPreload10", NestedPreload10)
	measureAndRun(t, "84 NestedPreload10", OldNestedPreload10)

	measureAndRun(t, "85 NestedPreload11", NestedPreload11)
	measureAndRun(t, "85 NestedPreload11", OldNestedPreload11)

	measureAndRun(t, "86 NestedPreload12", NestedPreload12)
	measureAndRun(t, "86 NestedPreload12", OldNestedPreload12)

	measureAndRun(t, "87 ManyToManyPreloadWithMultiPrimaryKeys", ManyToManyPreloadWithMultiPrimaryKeys)
	measureAndRun(t, "87 ManyToManyPreloadWithMultiPrimaryKeys", OldManyToManyPreloadWithMultiPrimaryKeys)

	measureAndRun(t, "88 ManyToManyPreloadForNestedPointer", ManyToManyPreloadForNestedPointer)
	measureAndRun(t, "88 ManyToManyPreloadForNestedPointer", OldManyToManyPreloadForNestedPointer)

	measureAndRun(t, "89 NestedManyToManyPreload", NestedManyToManyPreload)
	measureAndRun(t, "89 NestedManyToManyPreload", OldNestedManyToManyPreload)

	measureAndRun(t, "90 NestedManyToManyPreload2", NestedManyToManyPreload2)
	measureAndRun(t, "90 NestedManyToManyPreload2", OldNestedManyToManyPreload2)

	measureAndRun(t, "91 NestedManyToManyPreload3", NestedManyToManyPreload3)
	measureAndRun(t, "91 NestedManyToManyPreload3", OldNestedManyToManyPreload3)

	measureAndRun(t, "92 NestedManyToManyPreload3ForStruct", NestedManyToManyPreload3ForStruct)
	measureAndRun(t, "92 NestedManyToManyPreload3ForStruct", OldNestedManyToManyPreload3ForStruct)

	measureAndRun(t, "93 NestedManyToManyPreload4", NestedManyToManyPreload4)
	measureAndRun(t, "93 NestedManyToManyPreload4", OldNestedManyToManyPreload4)

	measureAndRun(t, "94 ManyToManyPreloadForPointer", ManyToManyPreloadForPointer)
	measureAndRun(t, "94 ManyToManyPreloadForPointer", OldManyToManyPreloadForPointer)

	measureAndRun(t, "95 NilPointerSlice", NilPointerSlice)
	measureAndRun(t, "95 NilPointerSlice", OldNilPointerSlice)

	measureAndRun(t, "96 NilPointerSlice2", NilPointerSlice2)
	measureAndRun(t, "96 NilPointerSlice2", OldNilPointerSlice2)

	measureAndRun(t, "97 PrefixedPreloadDuplication", PrefixedPreloadDuplication)
	measureAndRun(t, "97 PrefixedPreloadDuplication", OldPrefixedPreloadDuplication)

	measureAndRun(t, "98 FirstAndLast", FirstAndLast)
	measureAndRun(t, "98 FirstAndLast", OldFirstAndLast)

	measureAndRun(t, "99 FirstAndLastWithNoStdPrimaryKey", FirstAndLastWithNoStdPrimaryKey)
	measureAndRun(t, "99 FirstAndLastWithNoStdPrimaryKey", OldFirstAndLastWithNoStdPrimaryKey)

	measureAndRun(t, "100 UIntPrimaryKey", UIntPrimaryKey)
	measureAndRun(t, "100 UIntPrimaryKey", OldUIntPrimaryKey)

	measureAndRun(t, "101 StringPrimaryKeyForNumericValueStartingWithZero", StringPrimaryKeyForNumericValueStartingWithZero)
	measureAndRun(t, "101 StringPrimaryKeyForNumericValueStartingWithZero", OldStringPrimaryKeyForNumericValueStartingWithZero)

	measureAndRun(t, "102 FindAsSliceOfPointers", FindAsSliceOfPointers)
	measureAndRun(t, "102 FindAsSliceOfPointers", OldFindAsSliceOfPointers)

	measureAndRun(t, "103 SearchWithPlainSQL", SearchWithPlainSQL)
	measureAndRun(t, "103 SearchWithPlainSQL", OldSearchWithPlainSQL)

	measureAndRun(t, "104 SearchWithStruct", SearchWithStruct)
	measureAndRun(t, "104 SearchWithStruct", OldSearchWithStruct)

	measureAndRun(t, "105 SearchWithMap", SearchWithMap)
	measureAndRun(t, "105 SearchWithMap", OldSearchWithMap)

	measureAndRun(t, "106 SearchWithEmptyChain", SearchWithEmptyChain)
	measureAndRun(t, "106 SearchWithEmptyChain", OldSearchWithEmptyChain)

	measureAndRun(t, "107 Select", Select)
	measureAndRun(t, "107 Select", OldSelect)

	measureAndRun(t, "108 OrderAndPluck", OrderAndPluck)
	measureAndRun(t, "108 OrderAndPluck", OldOrderAndPluck)

	measureAndRun(t, "109 Limit", Limit)
	measureAndRun(t, "109 Limit", OldLimit)

	measureAndRun(t, "110 Offset", Offset)
	measureAndRun(t, "110 Offset", OldOffset)

	measureAndRun(t, "111 Or", Or)
	measureAndRun(t, "111 Or", OldOr)

	measureAndRun(t, "112 Count", Count)
	measureAndRun(t, "112 Count", OldCount)

	measureAndRun(t, "113 Not", Not)
	measureAndRun(t, "113 Not", OldNot)

	measureAndRun(t, "114 FillSmallerStruct", FillSmallerStruct)
	measureAndRun(t, "114 FillSmallerStruct", OldFillSmallerStruct)

	measureAndRun(t, "115 FindOrInitialize", FindOrInitialize)
	measureAndRun(t, "115 FindOrInitialize", OldFindOrInitialize)

	measureAndRun(t, "116 FindOrCreate", FindOrCreate)
	measureAndRun(t, "116 FindOrCreate", OldFindOrCreate)

	measureAndRun(t, "117 SelectWithEscapedFieldName", SelectWithEscapedFieldName)
	measureAndRun(t, "117 SelectWithEscapedFieldName", OldSelectWithEscapedFieldName)

	measureAndRun(t, "118 SelectWithVariables", SelectWithVariables)
	measureAndRun(t, "118 SelectWithVariables", OldSelectWithVariables)

	measureAndRun(t, "119 FirstAndLastWithRaw", FirstAndLastWithRaw)
	measureAndRun(t, "119 FirstAndLastWithRaw", OldFirstAndLastWithRaw)

	measureAndRun(t, "120 ScannableSlices", ScannableSlices)
	measureAndRun(t, "120 ScannableSlices", OldScannableSlices)

	measureAndRun(t, "121 Scopes", Scopes)
	measureAndRun(t, "121 Scopes", OldScopes)

	measureAndRun(t, "122 Update", Update)
	measureAndRun(t, "122 Update", OldUpdate)

	measureAndRun(t, "123 UpdateWithNoStdPrimaryKeyAndDefaultValues", UpdateWithNoStdPrimaryKeyAndDefaultValues)
	measureAndRun(t, "123 UpdateWithNoStdPrimaryKeyAndDefaultValues", OldUpdateWithNoStdPrimaryKeyAndDefaultValues)

	measureAndRun(t, "124 Updates", Updates)
	measureAndRun(t, "124 Updates", OldUpdates)

	measureAndRun(t, "125 UpdateColumn", UpdateColumn)
	measureAndRun(t, "125 UpdateColumn", OldUpdateColumn)

	measureAndRun(t, "126 SelectWithUpdate", SelectWithUpdate)
	measureAndRun(t, "126 SelectWithUpdate", OldSelectWithUpdate)

	measureAndRun(t, "127 SelectWithUpdateWithMap", SelectWithUpdateWithMap)
	measureAndRun(t, "127 SelectWithUpdateWithMap", OldSelectWithUpdateWithMap)

	measureAndRun(t, "128 OmitWithUpdate", OmitWithUpdate)
	measureAndRun(t, "128 OmitWithUpdate", OldOmitWithUpdate)

	measureAndRun(t, "129 OmitWithUpdateWithMap", OmitWithUpdateWithMap)
	measureAndRun(t, "129 OmitWithUpdateWithMap", OldOmitWithUpdateWithMap)

	measureAndRun(t, "130 SelectWithUpdateColumn", SelectWithUpdateColumn)
	measureAndRun(t, "130 SelectWithUpdateColumn", OldSelectWithUpdateColumn)

	measureAndRun(t, "131 OmitWithUpdateColumn", OmitWithUpdateColumn)
	measureAndRun(t, "131 OmitWithUpdateColumn", OldOmitWithUpdateColumn)

	measureAndRun(t, "132 UpdateColumnsSkipsAssociations", UpdateColumnsSkipsAssociations)
	measureAndRun(t, "132 UpdateColumnsSkipsAssociations", OldUpdateColumnsSkipsAssociations)

	measureAndRun(t, "133 UpdatesWithBlankValues", UpdatesWithBlankValues)
	measureAndRun(t, "133 UpdatesWithBlankValues", OldUpdatesWithBlankValues)

	measureAndRun(t, "134 UpdatesTableWithIgnoredValues", UpdatesTableWithIgnoredValues)
	measureAndRun(t, "134 UpdatesTableWithIgnoredValues", OldUpdatesTableWithIgnoredValues)

	measureAndRun(t, "135 UpdateDecodeVirtualAttributes", UpdateDecodeVirtualAttributes)
	measureAndRun(t, "135 UpdateDecodeVirtualAttributes", OldUpdateDecodeVirtualAttributes)

	measureAndRun(t, "136 ToDBNameGenerateFriendlyName", ToDBNameGenerateFriendlyName)
	measureAndRun(t, "136 ToDBNameGenerateFriendlyName", OldToDBNameGenerateFriendlyName)

	measureAndRun(t, "137 SkipSaveAssociation", SkipSaveAssociation)
	measureAndRun(t, "137 SkipSaveAssociation", OldSkipSaveAssociation)

	totalsNew := &Measure{
		netAllocs: 0,
		netBytes:  0,
		duration:  0,
	}
	totalsOld := &Measure{
		netAllocs: 0,
		netBytes:  0,
		duration:  0,
	}

	table := "\n| Test name | Allocs | Bytes | Duration  | Dif Allocs | Dif Bytes | Dif Duration |\n"
	table += "| :-------: | -----: | ----: | --------: | ---------: | --------: | -----------: |\n"
	for _, meas := range measuresData {
		if !meas.isNew {
			continue
		}
		totalsNew.add(meas)

		table += fmt.Sprintf("| *%s* | %d | %d | %s | | | |\n", meas.name, meas.netAllocs, meas.netBytes, DurationToString(meas.duration))

		//t.Logf("[1] %s : %d allocs %d bytes %s", meas.name, meas.netAllocs, meas.netBytes, DurationToString(meas.duration))

		if meas.pair == nil {
			t.Logf("ERROR : measurement with id %d has NO PAIR!", meas.pairId)
		} else {
			totalsOld.add(meas.pair)
			table += fmt.Sprintf("| %s | %d | %d | %s | | | |\n", meas.pair.name, meas.pair.netAllocs, meas.pair.netBytes, DurationToString(meas.pair.duration))

			table += "| diffs | | | |"

			difAllocs := int64(meas.pair.netAllocs - meas.netAllocs)
			if difAllocs == 0 {
				table += " :zzz: |"
			} else if difAllocs > 0 {
				table += fmt.Sprintf(" :zap: %d |", difAllocs)
			} else {
				table += fmt.Sprintf(" :snail: %d |", -difAllocs)
			}

			difBytes := int64(meas.pair.netBytes - meas.netBytes)
			if difBytes == 0 {
				table += " :zzz: |"
			} else if difBytes > 0 {
				table += fmt.Sprintf(" :zap: %d |", difBytes)
			} else {
				table += fmt.Sprintf(" :snail: %d |", -difBytes)
			}

			difDuration := int64(meas.pair.duration - meas.duration)
			if difDuration == 0 {
				table += " :zzz: |"
			} else if difDuration > 0 {
				table += fmt.Sprintf(" :zap: %s |", DurationToString(uint64(difDuration)))
			} else {
				table += fmt.Sprintf(" :snail: %s |", DurationToString(uint64(-difDuration)))
			}
			table += "\n"

			//t.Logf("[2] %s : %d allocs %d bytes %s", meas.pair.name, meas.pair.netAllocs, meas.pair.netBytes, DurationToString(meas.pair.duration))
			t.Logf("[%s] Diffs : %d allocs %d bytes %d nanoseconds", meas.name, difAllocs, difBytes, difDuration)
		}
	}

	table += fmt.Sprintf("| TOTAL (original) | %d | %d | %s | | | |\n", totalsOld.netAllocs, totalsOld.netBytes, DurationToString(totalsOld.duration))
	table += fmt.Sprintf("| TOTAL (new) | %d | %d | %s | | | |\n", totalsNew.netAllocs, totalsNew.netBytes, DurationToString(totalsNew.duration))
	t.Log(table)

}
func (m *Measure) add(from *Measure) {
	m.duration += from.duration
	m.netAllocs += from.netAllocs
	m.netBytes += from.netBytes
}

//Copied from time.Duration
func DurationToString(u uint64) string {
	// Largest time is 2540400h10m10.000000000s
	var buf [32]byte
	w := len(buf)

	if u < uint64(time.Second) {
		// Special case: if duration is smaller than a second,
		// use smaller units, like 1.2ms
		var prec int
		w--
		buf[w] = 's'
		w--
		switch {
		case u == 0:
			return "nothing."
		case u < uint64(time.Microsecond):
			// print nanoseconds
			prec = 0
			buf[w] = 'n'
		case u < uint64(time.Millisecond):
			// print microseconds
			prec = 3
			// U+00B5 'µ' micro sign == 0xC2 0xB5
			w-- // Need room for two bytes.
			copy(buf[w:], "µ")
		default:
			// print milliseconds
			prec = 6
			buf[w] = 'm'
		}
		w, u = fmtFrac(buf[:w], u, prec)
		w = fmtInt(buf[:w], u)
	} else {
		w--
		buf[w] = 's'

		w, u = fmtFrac(buf[:w], u, 9)

		// u is now integer seconds
		w = fmtInt(buf[:w], u%60)
		u /= 60

		// u is now integer minutes
		if u > 0 {
			w--
			buf[w] = 'm'
			w = fmtInt(buf[:w], u%60)
			u /= 60

			// u is now integer hours
			// Stop at hours because days can be different lengths.
			if u > 0 {
				w--
				buf[w] = 'h'
				w = fmtInt(buf[:w], u)
			}
		}
	}

	return string(buf[w:])
}

//Copied from time.Duration
// fmtFrac formats the fraction of v/10**prec (e.g., ".12345") into the
// tail of buf, omitting trailing zeros.  it omits the decimal
// point too when the fraction is 0.  It returns the index where the
// output bytes begin and the value v/10**prec.
func fmtFrac(buf []byte, v uint64, prec int) (nw int, nv uint64) {
	// Omit trailing zeros up to and including decimal point.
	w := len(buf)
	print := false
	for i := 0; i < prec; i++ {
		digit := v % 10
		print = print || digit != 0
		if print {
			w--
			buf[w] = byte(digit) + '0'
		}
		v /= 10
	}
	if print {
		w--
		buf[w] = '.'
	}
	return w, v
}

//Copied from time.Duration
// fmtInt formats v into the tail of buf.
// It returns the index where the output begins.
func fmtInt(buf []byte, v uint64) int {
	w := len(buf)
	if v == 0 {
		w--
		buf[w] = '0'
	} else {
		for v > 0 {
			w--
			buf[w] = byte(v%10) + '0'
			v /= 10
		}
	}
	return w
}
