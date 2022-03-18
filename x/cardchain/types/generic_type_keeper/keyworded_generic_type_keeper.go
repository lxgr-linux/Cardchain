package generic_type_keeper

import(
  sdk "github.com/cosmos/cosmos-sdk/types"
  "github.com/cosmos/cosmos-sdk/codec"
)

type KeywordedGenericTypeKeeper[T codec.ProtoMarshaler] struct {
  GenericTypeKeeper[T]
  KeyWords []string
}

// NewKGTK Returns a new KeywordedGenericTypeKeeper
func NewKGTK[T codec.ProtoMarshaler](key sdk.StoreKey, cdc codec.BinaryCodec, getEmpty func() T, keywords []string) KeywordedGenericTypeKeeper[T] {
  gtk := KeywordedGenericTypeKeeper[T] {
    NewGTK[T](key, cdc, getEmpty),
    keywords,
  }
  return gtk
}

// Get Gets an object from store
func (gtk KeywordedGenericTypeKeeper[T]) Get(ctx sdk.Context, keyword string) (T) {
  store := ctx.KVStore(gtk.Key)
	bz := store.Get([]byte(keyword))

  gotten := gtk.getEmpty()
	gtk.cdc.MustUnmarshal(bz, gotten)
	return gotten
}

// Set Sets an object in store
func (gtk KeywordedGenericTypeKeeper[T]) Set(ctx sdk.Context, keyword string, new T) {
	store := ctx.KVStore(gtk.Key)
	store.Set([]byte(keyword), gtk.cdc.MustMarshal(new))
}

// GetAll Gets all objs from store
func (gtk KeywordedGenericTypeKeeper[T]) GetAll(ctx sdk.Context) (all []T) {
	for _, keyWord := range gtk.KeyWords {
		all = append(all, gtk.Get(ctx, keyWord))
	}
	return
}

// GetNumber Gets the number of all objs in store
func (gtk KeywordedGenericTypeKeeper[T]) GetNumber(ctx sdk.Context) (id uint64) {
  return uint64(len(gtk.KeyWords))
}
