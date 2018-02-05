#pragma once
#include <eoslib/types.hpp>
#include <eoslib/datastream.hpp>
#include <eoslib/raw_fwd.hpp>

namespace eosio { namespace raw {
   
   template<typename Stream> inline void pack( Stream& s, const user_Info& value ) {
      raw::pack(s, value.user_type);
      raw::pack(s, value.email);
      raw::pack(s, value.role_type);
      raw::pack(s, value.signature);
   }
   template<typename Stream> inline void unpack( Stream& s, user_Info& value ) {
      raw::unpack(s, value.user_type);
      raw::unpack(s, value.email);
      raw::unpack(s, value.role_type);
      raw::unpack(s, value.signature);
   }
   template<typename Stream> inline void pack( Stream& s, const add_user_req& value ) {
      raw::pack(s, value.user_name);
      raw::pack(s, value.info);
   }
   template<typename Stream> inline void unpack( Stream& s, add_user_req& value ) {
      raw::unpack(s, value.user_name);
      raw::unpack(s, value.info);
   }
} }

#include <eoslib/raw.hpp>
namespace eosio {
   void print_ident(int n){while(n-->0){print("  ");}};
   template<typename Type>
   Type current_message_ex() {
      uint32_t size = message_size();
      char* data = (char *)eosio::malloc(size);
      assert(data && read_message(data, size) == size, "error reading message");
      Type value;
      eosio::raw::unpack(data, size, value);
      eosio::free(data);
      return value;
   }


   template<>
   user_Info current_message<user_Info>() {
      return current_message_ex<user_Info>();
   }

   template<>
   add_user_req current_message<add_user_req>() {
      return current_message_ex<add_user_req>();
   }
} //eosio

