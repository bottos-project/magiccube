#pragma once
#include <eoslib/types.hpp>
#include <eoslib/datastream.hpp>
#include <eoslib/raw_fwd.hpp>

namespace eosio { namespace raw {
   
   template<typename Stream> inline void pack( Stream& s, const user_base_Info& value ) {
      raw::pack(s, value.user_type);
      raw::pack(s, value.email);
      raw::pack(s, value.role_type);
      raw::pack(s, value.signature);
      raw::pack(s, value.active_key);
      raw::pack(s, value.owner_key);
   }
   template<typename Stream> inline void unpack( Stream& s, user_base_Info& value ) {
      raw::unpack(s, value.user_type);
      raw::unpack(s, value.email);
      raw::unpack(s, value.role_type);
      raw::unpack(s, value.signature);
      raw::unpack(s, value.active_key);
      raw::unpack(s, value.owner_key);
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

   void dump(const user_base_Info& value, int tab=0) {
      print_ident(tab);print("user_type:[");printi(uint64_t(value.user_type));print("]\n");
      print_ident(tab);print("email:[");prints(value.email.get_data());print("]\n");
      print_ident(tab);print("role_type:[");printi(value.role_type);print("]\n");
      print_ident(tab);print("signature:[");prints(value.signature.get_data());print("]\n");
      print_ident(tab);print("active_key:[");prints(value.active_key.get_data());print("]\n");
      print_ident(tab);print("owner_key:[");prints(value.owner_key.get_data());print("]\n");      
   }
   template<>
   user_base_Info current_message<user_base_Info>() {
      return current_message_ex<user_base_Info>();
   }
   void dump(const add_user_req& value, int tab=0) {
      print_ident(tab);print("user_name:[");prints(value.user_name.get_data());print("]\n");
      print_ident(tab);print("info:[");print("\n"); eosio::dump(value.info, tab+1);print_ident(tab);print("]\n");
   }
   template<>
   add_user_req current_message<add_user_req>() {
      return current_message_ex<add_user_req>();
   }
} //eosio

