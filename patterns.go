package main

var hwss = []cellCoord{
	{4, 4}, {4, 5}, {4, 6}, {4, 7}, {4, 8}, {4, 9},
	{5, 3}, {5, 9},
	{6, 9},
	{7, 3}, {7, 8},
	{8, 5}, {8, 6},
}

var pulsar = []cellCoord{
	{3, 5}, {3, 6}, {3, 7}, {3, 11}, {3, 12}, {3, 13},
	{5, 3}, {5, 8}, {5, 10}, {5, 15},
	{6, 3}, {6, 8}, {6, 10}, {6, 15},
	{7, 3}, {7, 8}, {7, 10}, {7, 15},
	{8, 5}, {8, 6}, {8, 7}, {8, 11}, {8, 12}, {8, 13},

	{10, 5}, {10, 6}, {10, 7}, {10, 11}, {10, 12}, {10, 13},
	{11, 3}, {11, 8}, {11, 10}, {11, 15},
	{12, 3}, {12, 8}, {12, 10}, {12, 15},
	{13, 3}, {13, 8}, {13, 10}, {13, 15},
	{15, 5}, {15, 6}, {15, 7}, {15, 11}, {15, 12}, {15, 13},
}

// GospersGliderGun is a Gosper's Glider Gun
var gospersGliderGun = []cellCoord{
	{2, 26},
	{3, 24}, {3, 26},
	{4, 14}, {4, 15}, {4, 22}, {4, 23}, {4, 36}, {4, 37},
	{5, 13}, {5, 17}, {5, 22}, {5, 23}, {5, 36}, {5, 37},
	{6, 2}, {6, 3}, {6, 12}, {6, 18}, {6, 22}, {6, 23},
	{7, 2}, {7, 3}, {7, 12}, {7, 16}, {7, 18}, {7, 19}, {7, 24}, {7, 26},
	{8, 12}, {8, 18}, {8, 26},
	{9, 13}, {9, 17},
	{10, 14}, {10, 15},
}