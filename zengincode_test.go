package zengincode

import (
	"os"
	"testing"
)

func TestBankNew(t *testing.T) {
	os.Setenv("ZENGIN_SOURCE_INCLUDE", "TRUE")
	tests := []struct {
		name    string
		code    string
		want    *Bank
		wantErr bool
	}{
		{
			name: "ufj",
			code: "0005",
			want: &Bank{
				Name: "三菱ＵＦＪ",
				Kana: "ミツビシユ－エフジエイ",
				Hira: "みつびしゆ－えふじえい",
				Roma: "mitsubishiyu-efujiei",
			},
			wantErr: false,
		},
		{
			name: "mizuho",
			code: "0001",
			want: &Bank{
				Name: "みずほ",
				Kana: "ミズホ",
				Hira: "みずほ",
				Roma: "mizuho",
			},
			wantErr: false,
		},
		{
			name: "mitsuisumitomo",
			code: "0009",
			want: &Bank{
				Name: "三井住友",
				Kana: "ミツイスミトモ",
				Hira: "みついすみとも",
				Roma: "mitsuisumitomo",
			},
			wantErr: false,
		},
		{
			name: "suruga",
			code: "0150",
			want: &Bank{
				Name: "スルガ",
				Kana: "スルガ",
				Hira: "するが",
				Roma: "suruga",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			bank := got[tt.code]
			if tt.want.Name != bank.Name {
				t.Errorf("New() Name = %s, want %s", bank.Name, tt.want.Name)
				return
			}
			if tt.want.Kana != bank.Kana {
				t.Errorf("New() Kana = %s, want %s", bank.Kana, tt.want.Kana)
				return
			}
			if tt.want.Hira != bank.Hira {
				t.Errorf("New() Hira = %s, want %s", bank.Hira, tt.want.Hira)
				return
			}
			if tt.want.Roma != bank.Roma {
				t.Errorf("New() Roma = %s, want %s", bank.Roma, tt.want.Roma)
				return
			}
		})
	}
}

func TestBranch(t *testing.T) {
	os.Setenv("ZENGIN_SOURCE_INCLUDE", "TRUE")
	tests := []struct {
		name       string
		code       string
		branchCode string
		want       *Branch
		wantErr    bool
	}{
		{
			name:       "ufj roppongi",
			code:       "0005",
			branchCode: "045",
			want: &Branch{
				Code: "045",
				Name: "六本木",
				Kana: "ロツポンギ",
				Hira: "ろつぽんぎ",
				Roma: "ropponngi",
			},
			wantErr: false,
		},
		{
			name:       "mizuho ebisu",
			code:       "0001",
			branchCode: "188",
			want: &Branch{
				Name: "恵比寿",
				Kana: "エビス",
				Hira: "えびす",
				Roma: "ebisu",
			},
			wantErr: false,
		},
		{
			name:       "mitsuisumitomo gifu",
			code:       "0009",
			branchCode: "407",
			want: &Branch{
				Name: "岐阜",
				Kana: "ギフ",
				Hira: "ぎふ",
				Roma: "gifu",
			},
			wantErr: false,
		},
		{
			name:       "suruga",
			code:       "0150",
			branchCode: "640",
			want: &Branch{
				Name: "本店",
				Kana: "ホンテン",
				Hira: "ほんてん",
				Roma: "honten",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			branch := got[tt.code].Branches[tt.branchCode]
			if tt.want.Name != branch.Name {
				t.Errorf("New() Name = %s, want %s", branch.Name, tt.want.Name)
				return
			}
			if tt.want.Kana != branch.Kana {
				t.Errorf("New() Kana = %s, want %s", branch.Kana, tt.want.Kana)
				return
			}
			if tt.want.Hira != branch.Hira {
				t.Errorf("New() Hira = %s, want %s", branch.Hira, tt.want.Hira)
				return
			}
			if tt.want.Roma != branch.Roma {
				t.Errorf("New() Roma = %s, want %s", branch.Roma, tt.want.Roma)
				return
			}
		})
	}
}
